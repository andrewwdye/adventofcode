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

struct Machine {
    a: (i32, i32),
    b: (i32, i32),
    prize: (i32, i32)
}

impl Machine {
    fn new(lines: &mut std::str::Lines) -> Self {
        let re = regex::Regex::new(r"\d+").unwrap();
        let a: Vec<i32> = re.find_iter(lines.next().unwrap()).map(|m| m.as_str().parse::<i32>().unwrap()).collect();
        let b: Vec<i32> = re.find_iter(lines.next().unwrap()).map(|m| m.as_str().parse::<i32>().unwrap()).collect();
        let prize: Vec<i32> = re.find_iter(lines.next().unwrap()).map(|m| m.as_str().parse::<i32>().unwrap()).collect();
        _ = lines.next();
        Machine { a: (a[0], a[1]), b: (b[0], b[1]), prize: (prize[0], prize[1]) }
    }
}

fn solve1(input: &str) -> Result<i32, std::io::Error>{
    let mut lines = input.lines();
    let mut machines = Vec::new();
    while lines.clone().peekable().peek().is_some() {
        let m = Machine::new(&mut lines);
        machines.push(m);
    }
    let mut sum = 0;
    for m in machines {
        for a_presses in 0..=100 {
            let ax = m.a.0 * a_presses;
            let ay = m.a.1 * a_presses;
            let bx = m.prize.0 - ax;
            if bx % m.b.0 != 0 {
                continue;
            }
            let b_presses = bx / m.b.0;
            if ay + m.b.1 * b_presses != m.prize.1 {
                continue;
            }
            let tokens = 3 * a_presses + b_presses;
            sum += tokens;
        }
    }
    Ok(sum)
}

fn solve2(_: &str) -> Result<i32, std::io::Error>{
    unimplemented!()
}

#[cfg(test)]
mod tests {
    use super::*;

    const SAMPLE_INPUT: &str = "Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279";

    #[test]
    fn test_part1() {
        assert_eq!(solve1(SAMPLE_INPUT).unwrap(), 480);
    }
}
