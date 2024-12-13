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

const FREE_SPACE: i32 = -1;

#[derive(Clone, Copy)]
struct Segment {
    id: i32,
    len: i32
}

struct Disk {
    segments: Vec<Segment>
}

impl Disk {
    fn new(input: &[u8]) -> Disk {
        let mut segments = Vec::new();
        for i in 0..input.len() {
            let id = if i % 2 == 0 { i as i32 / 2 } else { FREE_SPACE };
            segments.push(Segment{id, len: (input[i] - b'0') as i32});
        }
        Disk { segments }
    }

    fn blocks(&self) -> Vec<i32> {
        let mut blocks = Vec::new();
        for s in &self.segments {
            for _ in 0..s.len {
                blocks.push(s.id);
            }
        }
        blocks
    }

    fn checksum(&self) -> i64 {
        let mut checksum = 0;
        for (i, &block) in self.blocks().iter().enumerate() {
            if block == FREE_SPACE {
                continue;
            }
            checksum += i as i64 * block as i64;
        }
        checksum
    }

    fn defrag(&mut self) {
        let mut i = 0;
        while i < self.segments.len() {
            // println!("{:?}", self.blocks());
            if self.segments[i].id == FREE_SPACE {
                let mut remaining = self.segments[i].len;
                let mut right = self.segments.len() - 1;
                while remaining > 0 && right > i {
                    // println!("{} {}", remaining, right);
                    if self.segments[right].id != FREE_SPACE && self.segments[right].len <= remaining {
                        let s = self.segments[right];
                        remaining -= self.segments[right].len;
                        self.segments[right].id = FREE_SPACE;
                        self.segments.insert(i, s);
                        i += 1;
                    }
                    right -= 1;
                }
                self.segments[i].len = remaining;
            }
            i += 1;
        }
    }
}

fn solve2(input: &str) -> Result<i64, std::io::Error>{
    let mut disk = Disk::new(input.trim().as_bytes());
    disk.defrag();
    Ok(disk.checksum())
}

#[cfg(test)]
mod tests {
    use super::*;

    const SAMPLE_INPUT: &str = "2333133121414131402";

    #[test]
    fn test_part1() {
        assert_eq!(solve1(SAMPLE_INPUT).unwrap(), 1928);
    }

    #[test]
    fn test_part2() {
        assert_eq!(solve2(SAMPLE_INPUT).unwrap(), 2858);
        assert_eq!(solve2("1313165").unwrap(), 169);

        // 0 0 . . . 1 1 1 . . . 2 . . . 3 3 3 . 4 4 . 5 5 5 5 . 6 6 6 6 . 7 7 7 . 8 8 8 8 9 9 .10
        // 0 010 9 9 1 1 1 7 7 7 2 4 4 . 3 3 3 . . . . 5 5 5 5 . 6 6 6 6 . . . . . 8 8 8 8 . . . .
        assert_eq!(solve2("233313312141413140211").unwrap(), 2910);
    }
}
