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
    let input = input.trim().as_bytes();
    let mut blocks = Vec::new();
    let mut left: usize = 0;
    let mut right: usize = input.len() - 1;
    let mut right_consumed = 0;
'input:
    loop {
        // Consume file
        let mut file_len = input[left] - b'0';
        if left == right {
            file_len = file_len - right_consumed;
        }
        for _ in 0..file_len {
            blocks.push(left / 2); // id is half of index
        }
        if left == right{
            break 'input;
        }
        left += 1;

        // Consume free space
        let free_len = input[left] - b'0';
        for _ in 0..free_len {
            while right_consumed >= input[right] - b'0' {
                right_consumed = 0;
                right -= 2;
                if right < left {
                    break 'input;
                }
            }
            blocks.push(right / 2); // id is half of index
            right_consumed += 1;
        }
        left += 1;
    }
    let mut checksum = 0;
    for (i, &block) in blocks.iter().enumerate() {
        checksum += i as i64 * block as i64;
    }
    Ok(checksum)
}

fn solve2(_: &str) -> Result<i64, std::io::Error>{
    unimplemented!()
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    const SAMPLE_INPUT: &str = "2333133121414131402";

    #[test]
    fn test_part1() {
        assert_eq!(solve1(SAMPLE_INPUT).unwrap(), 1928);
    }
}
