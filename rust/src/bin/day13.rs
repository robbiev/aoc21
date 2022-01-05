use std::error::Error;
use std::collections::HashSet;

#[derive(Debug, PartialEq, Eq, Hash, Clone)]
struct Point {
    x: i32,
    y: i32,
}

#[derive(Debug)]
struct Paper {
	dots:  Vec<Point>,
	folds: Vec<Point>,
}

fn fold_x(x: i32, dot: &Point) -> Point {
    if dot.x < x {
        return dot.clone() // we fold left only
    }
    Point{
        x: dot.x - ((dot.x - x) * 2),
        y: dot.y,
    }
}

fn fold_y(y: i32, dot: &Point) -> Point {
    if dot.y < y {
        return dot.clone() // we fold left only
    }
    Point{
        x: dot.x,
        y: dot.y - ((dot.y - y) * 2),
    }
}

fn solve(paper: Paper, max_folds: usize) -> i32 {
    let mut dot_map = HashSet::<Point>::new();
    let mut bounds = Point{ x: 0, y: 0 };
    for dot in paper.dots {
        if dot.x > bounds.x {
            bounds.x = dot.x;
        }
        if dot.y > bounds.y {
            bounds.y = dot.y;
        }
        dot_map.insert(dot);
    }

    for fold in &paper.folds[..max_folds] {
        let fold_func: Box<dyn Fn(&Point) -> Point> = if fold.x != 0 {
            bounds.x = bounds.x - fold.x - 1;
            Box::new(|dot| fold_x(fold.x, dot))
        } else if fold.y != 0 {
            bounds.y = bounds.y - fold.y - 1;
            Box::new(|dot| fold_y(fold.y, dot))
        } else {
            Box::new(|dot| dot.clone())
        };

        let mut replacements = Vec::<Point>::new();
        for dot in &dot_map {
            let new_dot = fold_func(dot);
            replacements.push(new_dot);
        }
        for replace in replacements {
            dot_map.insert(replace);
        }
    }

    let mut dot_count: i32 = 0;
    for dot in &dot_map {
        if dot.x >= 0 && dot.x <= bounds.x && dot.y >= 0 && dot.y <= bounds.y {
            dot_count += 1;
        }
    }

    for y in 0..=bounds.y {
        for x in 0..=bounds.x {
            if dot_map.contains(&Point{ x, y }) {
                print!("{}", "#");
            } else {
                print!("{}", ".");
            }
        }
        println!();
    }

    dot_count
}

const INPUT: &str = include_str!("day13_input.txt");

fn parse_input(input_txt: &str) -> Result<Paper, Box<dyn Error>> {
    let mut input = Paper {
        dots: vec![],
        folds: vec![]
    };
    for line in input_txt.lines().map(str::split_whitespace) {
        match line.collect::<Vec<_>>()[..] {
            [_, _, fold] => {
                let mut split = fold.split('=');
                let xy = split.next().ok_or("")?;
                let amount = split.next().ok_or("")?.parse()?;
                match xy {
                    "x" => input.folds.push(Point{ x: amount, y: 0 }),
                    "y" => input.folds.push(Point{ x: 0, y: amount }),
                    _ => continue
                };
            },
            [dot] => {
                let mut split = dot.split(",");
                let x = split.next().ok_or("")?;
                let y = split.next().ok_or("")?;
                input.dots.push(Point{ x: x.parse()?, y: y.parse()? });
            },
            _ => continue
        }
    }

    Ok(input)
}

fn main() -> Result<(), Box<dyn Error>> {
    let input = parse_input(INPUT)?;
    let max_folds = input.folds.len();
    println!("{:?}", solve(input, max_folds));
    Ok(())
}

#[test]
fn test_fold_y() -> Result<(), Box<dyn Error>> {
    let actual = fold_y(7, &Point{ x: 0, y: 14 });
    let expected = Point { x: 0, y: 0 };
    assert_eq!(actual, expected);
    Ok(())
}

#[test]
fn test_part1() -> Result<(), Box<dyn Error>> {
    const EXAMPLE1: &str = r#"6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5"#;
    let input = parse_input(EXAMPLE1)?;
    let actual = solve(input, 1);
    assert_eq!(actual, 17);
    Ok(())
}
