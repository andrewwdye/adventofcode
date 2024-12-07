use clap::Parser;
use std::fs::read_to_string;

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

fn solve1(input: &str) -> Result<i64, std::io::Error>{
    let mut sum = 0;
    for line in input.lines() {
        let mut it = line.split(": ");
        let total = it.next().unwrap().parse::<i64>().unwrap();
        let nums: Vec<i64> = it.next().unwrap().split(" ").map(|s| s.parse::<i64>().unwrap()).collect();
        if check(&nums, total) {
            sum += total;
        }
    }
    Ok(sum)
}

fn check(nums: &[i64], total: i64) -> bool {
    if nums.len() == 1 {
        return nums[0] == total;
    }
    let last = nums[nums.len()-1];
    let remaining = &nums[0..nums.len()-1];
    // Check add
    if last > total {
        return false;
    }
    if check(remaining, total - last) {
        return true;
    }
    // Check multiply
    if total % last != 0 {
        return false;
    }
    if check(remaining, total / last) {
        return true;
    }
    false
}

fn solve2(_: &str) -> Result<i64, std::io::Error>{
    unimplemented!()
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    const SAMPLE_INPUT: &str = "190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20";

    #[test]
    fn test_part1() {
        assert_eq!(solve1(SAMPLE_INPUT).unwrap(), 3749);
    }
}
