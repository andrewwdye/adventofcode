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

fn solve1(input: &str) -> Result<i32, std::io::Error>{
    let lines: Vec<&str> = input.lines().collect();
    let mut count = 0;
    for i in 0..lines.len() {
        for j in 0..lines[i].len() {
            for (di, dj) in [(-1, 0), (1, 0), (0, -1), (0, 1), (-1, -1), (-1, 1), (1, -1), (1, 1)].iter() {
                count += search1(&lines, i as i32, j as i32, *di, *dj, "XMAS");
            }
        }
    }
    Ok(count)
}

fn search1(lines: &Vec<&str>, i: i32, j: i32, di: i32, dj: i32, word: &str) -> i32 {
    if word.len() == 0 {
        return 1;
    }
    if i < 0 || i >= lines.len() as i32 {
        return 0;
    }
    if j < 0 || j >= lines[i as usize].len() as i32 {
        return 0;
    }
    if lines[i as usize].as_bytes()[j as usize] != word.as_bytes()[0] {
        return 0;
    }
    search1(lines, i + di, j + dj, di, dj, &word[1..])
}

fn solve2(input: &str) -> Result<i32, std::io::Error>{
    let lines: Vec<&str> = input.lines().collect();
    let mut count = 0;
    for i in 1..lines.len()-1 {
        for j in 1..lines[i].len()-1 {
            if find_xmas(&lines, i, j) {
                count += 1;
            }
        }
    }
    Ok(count)
}

fn find_xmas(lines: &Vec<&str>, i: usize, j: usize) -> bool {
    if lines[i].as_bytes()[j] != 'A' as u8{
        return false;
    }
    let mut left = vec![lines[i-1].as_bytes()[j-1], lines[i+1].as_bytes()[j+1]];
    let mut right = vec![lines[i+1].as_bytes()[j-1], lines[i-1].as_bytes()[j+1]];
    left.sort();
    right.sort();
    left == "MS".as_bytes().to_vec() && right == "MS".as_bytes().to_vec()
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn test_part1() {
        let input = "MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX";
        assert_eq!(solve1(input).unwrap(), 18);
    }

    #[test]
    fn test_part2() {
        let input = "MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX";
        assert_eq!(solve2(input).unwrap(), 9);
    }
}
