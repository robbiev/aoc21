use std::error::Error;
use std::collections::HashSet;
use std::cmp::max;

#[derive(Debug, PartialEq, Eq, Hash, Clone)]
struct Point {
    x: i32,
    y: i32,
}

type Bounds = Point;

#[derive(Debug)]
struct Paper {
	dots:  Vec<Point>,
	folds: Vec<Fold>,
}

#[derive(Debug)]
enum Fold {
    Left(i32),
    Up(i32)
}

fn fold_index(target_index: i32, fold_index: i32) -> i32 {
    if target_index < fold_index {
        return target_index // only fold left/up
    }
    target_index - ((target_index - fold_index) * 2)
}

fn solve(paper: Paper, max_folds: usize) -> i32 {
    let mut dot_set = HashSet::<Point>::with_capacity(paper.dots.len());
    let mut bounds = Bounds{ x: 0, y: 0 };
    for dot in paper.dots {
        bounds.x = max(bounds.x, dot.x);
        bounds.y = max(bounds.y, dot.y);
        dot_set.insert(dot);
    }

    for fold in &paper.folds[..max_folds] {
        match fold {
            Fold::Left(x) => bounds.x -= x + 1,
            Fold::Up(y) => bounds.y -= y + 1,
        }

        let mut folded_dots = Vec::<Point>::new();
        for dot in &dot_set {
            let folded_dot = match fold {
                Fold::Left(x) => Point{ x: fold_index(dot.x, *x), y: dot.y },
                Fold::Up(y) => Point{ x: dot.x, y: fold_index(dot.y, *y) },
            };
            folded_dots.push(folded_dot);
        }
        for folded_dot in folded_dots {
            dot_set.insert(folded_dot);
        }
    }

    for y in 0..=bounds.y {
        for x in 0..=bounds.x {
            if dot_set.contains(&Point{ x, y }) {
                print!("{}", "#");
            } else {
                print!("{}", ".");
            }
        }
        println!();
    }

    dot_set.iter()
        .filter(|Point { x, y }| (0..=bounds.x).contains(x) && (0..=bounds.y).contains(y))
        .count() as i32
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
                    "x" => input.folds.push(Fold::Left(amount)),
                    "y" => input.folds.push(Fold::Up(amount)),
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
fn test_fold_index() -> Result<(), Box<dyn Error>> {
    let actual = fold_index(14, 7);
    let expected = 0;
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
