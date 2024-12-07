use clap::Parser;
use std::{collections::{HashMap, HashSet}, fs::read_to_string, str::Lines};

#[derive(Parser)]
struct Cli {
    #[arg(short, long, value_parser = clap::value_parser!(u8).range(1..=2), default_value_t = 1)]
    part: u8,

    input: String
}

fn main() -> Result<(), std::io::Error>{
    let cli = Cli::parse();
    let input = read_to_string(cli.input)?;
    let result = match cli.part {
        1 => solve1(input.as_str())?,
        2 => solve2(input.as_str())?,
        _ => unreachable!(),
    };
    println!("{result}");
    Ok(())
}

fn solve1(input: &str) -> Result<i32, std::io::Error>{
    let mut lines = input.lines();
    let rules = parse_rules(&mut lines);
    let updates = parse_updates(&mut lines);
    let mut middle_sums = 0;
    for update in updates.iter() {
        if is_valid(update, &rules) {
            middle_sums += update.get(update.len()/2).unwrap();
        }
    }
    Ok(middle_sums)
}

fn solve2(input: &str) -> Result<i32, std::io::Error>{
    let mut lines = input.lines();
    let rules = parse_rules(&mut lines);
    let updates = parse_updates(&mut lines);
    let mut fixed_middle_sums = 0;
    for update in updates.iter() {
        if !is_valid(update, &rules) {
            let update = fix(update, &rules);
            fixed_middle_sums += update.get(update.len()/2).unwrap();
        }
    }
    Ok(fixed_middle_sums)
}

/// Returns a map of rules where the key must come after the values in the set
fn parse_rules(lines: &mut Lines) -> HashMap<i32, HashSet<i32>> {
    let mut rules: HashMap<i32, HashSet<i32>> = HashMap::new();
    for line in lines {
        if line == "" {
            break;
        }
        let b: Vec<i32> = line.split('|').map(|s| s.parse().unwrap()).collect();
        rules.entry(b[1]).or_default().insert(b[0]);
    }
    rules
}

fn parse_updates(lines: &mut Lines) -> Vec<Vec<i32>> {
    let mut updates: Vec<Vec<i32>> = Vec::new();
    for line in lines {
        updates.push(line.split(',').map(|s| s.parse().unwrap()).collect());
    }
    updates
}

fn is_valid(update: &Vec<i32>, rules: &HashMap<i32, HashSet<i32>>) -> bool {
    for (i, left) in update.iter().enumerate() {
        for right in update.iter().skip(i+1) {
            if let Some(before) = rules.get(left) {
                if before.contains(right) {
                    return false;
                }
            }
        }
    }
    true
}

fn fix(update: &Vec<i32>, rules: &HashMap<i32, HashSet<i32>>) -> Vec<i32> {
    let mut fixed = update.clone();
    let mut done = 0;
    // Loop through remaining and find one that is allowed to be left of all others
'remaining:
    while done < fixed.len() {
    'find_left:
        for i in done..fixed.len() {
            let left = fixed[i];
            for j in done..fixed.len() {
                if i == j {
                    continue;
                }
                let right = fixed[j];
                if let Some(before) = rules.get(&left) {
                    if before.contains(&right) {
                        continue 'find_left;
                    }
                }
            }
            fixed.swap(i, done);
            done += 1;
            continue 'remaining;
        }
        unreachable!();
    }
    fixed
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    const SAMPLE_INPUT: &str = "47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47";

    #[test]
    fn test_part1() {
        assert_eq!(solve1(SAMPLE_INPUT).unwrap(), 143);
    }

    #[test]
    fn test_part2() {
        assert_eq!(solve2(SAMPLE_INPUT).unwrap(), 123);
    }
}
