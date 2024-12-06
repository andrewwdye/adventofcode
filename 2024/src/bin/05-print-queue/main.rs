use clap::Parser;
use std::{collections::{HashMap, HashSet}, fs::read_to_string};

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
    let mut must_come_after: HashMap<i32, HashSet<i32>> = HashMap::new();
    for line in &mut lines {
        if line == "" {
            break;
        }
        let b: Vec<i32> = line.split('|').map(|s| s.parse().unwrap()).collect();
        must_come_after.entry(b[1]).or_default().insert(b[0]);
    }
    // println!("Must come after: {:?}", must_come_after);
    let mut middle_sums = 0;
    'line_check:
    for line in lines {
        let pages: Vec<i32> = line.split(',').map(|s| s.parse().unwrap()).collect();
        // println!("Pages: {:?}", pages);
        for (i, left) in pages.iter().enumerate() {
            // println!("Checking {}", i);
            for right in pages.iter().skip(i+1) {
                // println!("Checking values {} -> {}", left, right);
                if let Some(before) = must_come_after.get(left) {
                    if before.contains(right) {
                        continue 'line_check;
                    }
                }
            }
        }
        middle_sums += pages.get(pages.len()/2).unwrap();
    }
    Ok(middle_sums)
}

fn solve2(_: &str) -> Result<i32, std::io::Error>{
    unimplemented!()
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn test_part1() {
        let input = "47|53
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
        assert_eq!(solve1(input).unwrap(), 0);
    }
}
