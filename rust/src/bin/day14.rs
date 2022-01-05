use std::collections::HashMap;
use std::error::Error;

#[derive(Debug)]
struct Polymerization {
    polymer_template: String,
    pair_insertion_rules: HashMap<(char, char), char>,
}

fn solve(polymerization: Polymerization, steps: usize) -> i64 {
    let polymer = Vec::from_iter(polymerization.polymer_template.chars());

    let mut pair_occurrences = HashMap::<(char, char), i64>::new();
    for pair in polymer[..].windows(2) {
        *pair_occurrences.entry((pair[0], pair[1])).or_default() += 1;
    }

    let mut frequencies: HashMap<char, i64> = HashMap::new();
    for poly in polymer {
        *frequencies.entry(poly).or_default() += 1;
    }

    for _ in 0..steps {
        let mut pair_occurrences_adjust = HashMap::<(char, char), i64>::new();
        for (pair, occurrences) in &pair_occurrences {
            if let Some(insert) = polymerization.pair_insertion_rules.get(pair) {
                // Inserting X in AB gives AXB
                // For every occurrence of AB:
                // - pair AX forms
                // - pair XB forms
                // - the original AB pair ceases to exist
                // - the frequency of X is increased
                *pair_occurrences_adjust.entry((pair.0, *insert)).or_default() += occurrences;
                *pair_occurrences_adjust.entry((*insert, pair.1)).or_default() += occurrences;
                *pair_occurrences_adjust.entry(pair.clone()).or_default() -= occurrences;
                *frequencies.entry(*insert).or_default() += occurrences;
            }
        }
        for (pair, occurrence_adjust) in pair_occurrences_adjust {
            *pair_occurrences.entry(pair).or_default() += occurrence_adjust;
        }
    }
    let (_, max) = frequencies.iter().max_by_key(|(_, v)| *v).unwrap();
    let (_, min) = frequencies.iter().min_by_key(|(_, v)| *v).unwrap();
    max - min
}

const INPUT: &str = include_str!("day14_input.txt");

fn parse_input(input_txt: &str) -> Result<Polymerization, Box<dyn Error>> {
    let mut input = Polymerization {
        polymer_template: "".to_string(),
        pair_insertion_rules: HashMap::new(),
    };
    let mut lines = input_txt.lines();
    input.polymer_template = lines.next().ok_or("")?.into();
    lines.next(); // eat blank line
    for line in lines.map(str::split_whitespace) {
        match line.collect::<Vec<_>>()[..] {
            [from, _, to] => {
                let key = (from.chars().nth(0).unwrap_or_default(), from.chars().nth(1).unwrap_or_default());
                input.pair_insertion_rules.insert(key, to.chars().next().ok_or("")?)
            },
            _ => continue
        };
    }

    Ok(input)
}

fn main() -> Result<(), Box<dyn Error>> {
    let input = parse_input(INPUT)?;
    println!("{:?}", solve(input, 10)); // 3411
    let input = parse_input(INPUT)?;
    println!("{:?}", solve(input, 40)); // 7477815755570
    Ok(())
}

#[test]
fn test_part1() -> Result<(), Box<dyn Error>> {
    const EXAMPLE1: &str = r#"NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C"#;
    let input = parse_input(EXAMPLE1)?;
    // NCNBCHB
    // {"CN": 1, "NB": 1, "NC": 2, "CB": 1, "HB": 1, "NN": 1, "BC": 1, "CH": 1}
    let actual = solve(input, 10);
    assert_eq!(actual, 1588);
    Ok(())
}
#[test]
fn test_part2() -> Result<(), Box<dyn Error>> {
    const EXAMPLE1: &str = r#"NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C"#;
    let input = parse_input(EXAMPLE1)?;
    let actual = solve(input, 40);
    assert_eq!(actual, 2188189693529);
    Ok(())
}
