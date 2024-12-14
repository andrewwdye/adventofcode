use clap::Parser;
use std::{collections::HashMap, fs::read_to_string};

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

type Lookaside = HashMap<(String, u8), usize>;

fn solve1(input: &str) -> Result<usize, std::io::Error>{
    let nums: Vec<String> = input.split_whitespace().map(|s: &str| s.to_string()).collect();
    Ok(blink(&nums, 25, &mut HashMap::new()))
}

fn blink(nums: &Vec<String>, count: u8, lookaside: &mut Lookaside) -> usize {
    if count == 0 {
        return nums.len();
    }
    let mut sum = 0;
    for n in nums {
        if let Some(&result) = lookaside.get(&(n.clone(), count)) {
            sum += result;
            continue;
        }
        let mut v = Vec::new();
        // Apply rule
        if n == "0" {
            v.push("1".to_string());
        } else if n.len() % 2 == 0 {
            let (left, right) = n.split_at(n.len() / 2);
            v.push(left.parse::<usize>().unwrap().to_string());
            v.push(right.parse::<usize>().unwrap().to_string());
        } else {
            let value: usize = n.parse().unwrap();
            v.push((value * 2024).to_string());
        }

        // Recurse
        let result = blink(&v, count - 1, lookaside);
        lookaside.insert((n.clone(), count), result);
        sum += result;
    }
    sum
}

fn solve2(input: &str) -> Result<usize, std::io::Error>{
    let nums: Vec<String> = input.split_whitespace().map(|s: &str| s.to_string()).collect();
    Ok(blink(&nums, 75, &mut HashMap::new()))
}

#[cfg(test)]
mod tests {
    use super::*;

    const SAMPLE_INPUT: &str = "125 17";

    #[test]
    fn test_part1() {
        let input = SAMPLE_INPUT.split_whitespace().map(|s: &str| s.to_string()).collect();
        assert_eq!(blink(&input, 0, &mut HashMap::new()), 2);
        assert_eq!(blink(&input, 1, &mut HashMap::new()), 3);
        assert_eq!(blink(&input, 2, &mut HashMap::new()), 4);
        assert_eq!(blink(&input, 25, &mut HashMap::new()), 55312);
    }
}
