use clap::Parser;
use std::fs::read_to_string;

#[derive(Parser)]
struct Cli {
    #[arg(short, long, value_parser = clap::value_parser!(u8).range(1..=2), default_value_t = 1)]
    part: u8,

    input: String,
}

fn main() -> Result<(), std::io::Error> {
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

fn solve1(input: &str) -> Result<String, std::io::Error> {
    let mut lines = input.lines();
    let mut computer = Computer::new(&mut lines);
    let program = Program::new(&mut lines);
    println!("{:?}", computer);
    println!("{:?}", program);
    Ok(computer.run(&program).to_string())
}

fn solve2(_: &str) -> Result<String, std::io::Error> {
    unimplemented!()
}

#[derive(Debug)]
enum Opcode {
    /// The adv instruction (opcode 0) performs division. The numerator is the value in the A register.
    /// The denominator is found by raising 2 to the power of the instruction's combo operand.
    /// (So, an operand of 2 would divide A by 4 (2^2); an operand of 5 would divide A by 2^B.)
    /// The result of the division operation is truncated to an integer and then written to the A register.
    Adv,
    /// The bxl instruction (opcode 1) calculates the bitwise XOR of register B and the instruction's literal operand,
    /// then stores the result in register B.
    Bxl,
    /// The bst instruction (opcode 2) calculates the value of its combo operand modulo 8
    /// (thereby keeping only its lowest 3 bits), then writes that value to the B register.
    Bst,
    /// The jnz instruction (opcode 3) does nothing if the A register is 0. However, if the A register is not zero,
    /// it jumps by setting the instruction pointer to the value of its literal operand; if this instruction jumps,
    /// the instruction pointer is not increased by 2 after this instruction.
    Jnz,
    /// The bxc instruction (opcode 4) calculates the bitwise XOR of register B and register C, then stores the result
    /// in register B. (For legacy reasons, this instruction reads an operand but ignores it.)
    Bxc,
    /// The out instruction (opcode 5) calculates the value of its combo operand modulo 8, then outputs that value.
    /// (If a program outputs multiple values, they are separated by commas.)
    Out,
    /// The bdv instruction (opcode 6) works exactly like the adv instruction except that the result is stored in the
    /// B register. (The numerator is still read from the A register.)
    Bdv,
    /// The cdv instruction (opcode 7) works exactly like the adv instruction except that the result is stored in the
    /// C register. (The numerator is still read from the A register.)
    Cdv,
}

impl From<i32> for Opcode {
    fn from(value: i32) -> Self {
        match value {
            0 => Opcode::Adv,
            1 => Opcode::Bxl,
            2 => Opcode::Bst,
            3 => Opcode::Jnz,
            4 => Opcode::Bxc,
            5 => Opcode::Out,
            6 => Opcode::Bdv,
            7 => Opcode::Cdv,
            _ => unreachable!(),
        }
    }
}

#[derive(Debug)]
struct Computer {
    registers: [i32; 3],
    ip: usize,
}

const A: usize = 0;
const B: usize = 1;
const C: usize = 2;

impl Computer {
    fn new(lines: &mut std::str::Lines) -> Self {
        let mut registers = [0, 0, 0];
        for i in 0..3 {
            registers[i] = lines
                .next()
                .unwrap()
                .split(": ")
                .skip(1)
                .next()
                .unwrap()
                .parse::<i32>()
                .expect("")
        }
        _ = lines.next();
        Computer { registers, ip: 0 }
    }

    fn run(&mut self, program: &Program) -> String {
        self.ip = 0;
        let mut result: Vec<String> = Vec::new();
        while self.ip < program.instructions.len() - 1 {
            let instruction = program.instructions[self.ip].into();
            let operand = program.instructions[self.ip + 1];
            if let Some(value) = self.do_instruction(instruction, operand) {
                result.push(value.to_string());
            }
        }
        result.join(",")
    }

    fn do_instruction(&mut self, instruction: Opcode, operand: i32) -> Option<i32> {
        let literal_operand = operand;
        let combo_operand = match literal_operand {
            0 | 1 | 2 | 3 | 7 => literal_operand,
            4 => self.registers[A],
            5 => self.registers[B],
            6 => self.registers[C],
            _ => unreachable!(),
        };
        println!("ip:         {:?}", self.ip);
        println!("registers   {:?}", self.registers);
        println!("instruction {:?}", instruction);
        println!("operand     {:?}/{:?}", literal_operand, combo_operand);
        self.ip += 2;
        let mut out = None;
        const BASE_2: i32 = 2;
        match instruction {
            Opcode::Adv => {
                println!(
                    "operation   reg[A] = {:?}/2^{:?} = {:?}",
                    self.registers[A],
                    combo_operand,
                    self.registers[A] / BASE_2.pow(combo_operand as u32)
                );
                self.registers[A] = self.registers[A] / BASE_2.pow(combo_operand as u32)
            }
            Opcode::Bxl => {
                println!(
                    "operation   reg[B] = {:?} XOR {:?} = {:?}",
                    self.registers[B],
                    literal_operand,
                    self.registers[B] ^ literal_operand
                );
                self.registers[B] ^= literal_operand
            }
            Opcode::Bst => {
                println!(
                    "operation   reg[B] = {:?} % 8 = {:?}",
                    combo_operand,
                    combo_operand % 8
                );
                self.registers[B] = combo_operand % 8
            }
            Opcode::Jnz => {
                if self.registers[A] != 0 {
                    println!("operation   ip = {:?}", literal_operand);
                    self.ip = literal_operand as usize;
                } else {
                    println!("operation   nop");
                }
            }
            Opcode::Bxc => {
                println!(
                    "operation   reg[B] = {:?} XOR {:?} = {:?}",
                    self.registers[B],
                    self.registers[C],
                    self.registers[B] ^ self.registers[C]
                );
                self.registers[B] ^= self.registers[C]
            }
            Opcode::Out => {
                println!(
                    "operation   out << {:?} % 8 = {:?}",
                    combo_operand,
                    combo_operand % 8
                );
                out = Some(combo_operand % 8);
            }
            Opcode::Bdv => {
                println!(
                    "operation   reg[B] = {:?}/2^{:?} = {:?}",
                    self.registers[A],
                    combo_operand,
                    self.registers[A] / BASE_2.pow(combo_operand as u32)
                );
                self.registers[B] = self.registers[A] / BASE_2.pow(combo_operand as u32)
            }
            Opcode::Cdv => {
                println!(
                    "operation   reg[C] = {:?}/2^{:?} = {:?}",
                    self.registers[A],
                    combo_operand,
                    self.registers[A] / BASE_2.pow(combo_operand as u32)
                );
                self.registers[C] = self.registers[A] / BASE_2.pow(combo_operand as u32)
            }
        }
        println!("");
        out
    }
}

#[derive(Debug)]
struct Program {
    instructions: Vec<i32>,
}

impl Program {
    fn new(lines: &mut std::str::Lines) -> Self {
        let instructions: Vec<i32> = lines
            .next()
            .unwrap()
            .split(": ")
            .skip(1)
            .next()
            .unwrap()
            .split(",")
            .map(|s: &str| s.parse::<i32>().expect(""))
            .collect();
        Program { instructions }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    const SAMPLE_INPUT: &str = "Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0";

    #[test]
    fn test_part1() {
        assert_eq!(solve1(SAMPLE_INPUT).unwrap(), "4,6,3,5,6,3,5,2,1,0");
    }

    #[test]
    fn test_examples() {
        // If register C contains 9, the program 2,6 would set register B to 1.
        let mut c = Computer {
            registers: [0, 0, 9],
            ip: 0,
        };
        let p = Program {
            instructions: vec![2, 6],
        };
        _ = c.run(&p);
        assert_eq!(c.registers[B], 1);

        // If register A contains 10, the program 5,0,5,1,5,4 would output 0,1,2.
        let mut c = Computer {
            registers: [10, 0, 0],
            ip: 0,
        };
        let p = Program {
            instructions: vec![5, 0, 5, 1, 5, 4],
        };
        let result = c.run(&p);
        assert_eq!(result, "0,1,2");

        // If register A contains 2024, the program 0,1,5,4,3,0 would output 4,2,5,6,7,7,7,7,3,1,0 and leave 0 in register A.
        let mut c = Computer {
            registers: [2024, 0, 0],
            ip: 0,
        };
        let p = Program {
            instructions: vec![0, 1, 5, 4, 3, 0],
        };
        let result = c.run(&p);
        assert_eq!(result, "4,2,5,6,7,7,7,7,3,1,0");
        assert_eq!(c.registers[A], 0);

        // If register B contains 29, the program 1,7 would set register B to 26.
        let mut c = Computer {
            registers: [0, 29, 0],
            ip: 0,
        };
        let p = Program {
            instructions: vec![1, 7],
        };
        _ = c.run(&p);
        assert_eq!(c.registers[B], 26);

        // If register B contains 2024 and register C contains 43690, the program 4,0 would set register B to 44354.
        let mut c = Computer {
            registers: [0, 2024, 43690],
            ip: 0,
        };
        let p = Program {
            instructions: vec![4, 0],
        };
        _ = c.run(&p);
        assert_eq!(c.registers[B], 44354);
    }

    #[test]
    fn test_adv() {
        let mut c = Computer {
            registers: [201, 1, 0],
            ip: 0,
        };
        _ = c.do_instruction(Opcode::Adv, 2);
        assert_eq!(c.registers, [50, 1, 0]);

        let mut c = Computer {
            registers: [201, 1, 0],
            ip: 0,
        };
        _ = c.do_instruction(Opcode::Adv, 5);
        assert_eq!(c.registers, [100, 1, 0]);
    }

    #[test]
    fn test_bxl() {
        let mut c = Computer {
            registers: [0, 0b100, 0],
            ip: 0,
        };
        _ = c.do_instruction(Opcode::Bxl, 0b11);
        assert_eq!(c.registers[B], 0b111);
    }

    #[test]
    fn test_bst() {
        let mut c = Computer {
            registers: [0, 15, 0],
            ip: 0,
        };
        _ = c.do_instruction(Opcode::Bst, 2);
        assert_eq!(c.registers[B], 2);

        let mut c = Computer {
            registers: [0, 15, 0],
            ip: 0,
        };
        _ = c.do_instruction(Opcode::Bst, 5);
        assert_eq!(c.registers[B], 7);
    }

    #[test]
    fn test_jnz() {
        let mut c = Computer {
            registers: [0, 0, 0],
            ip: 0,
        };
        _ = c.do_instruction(Opcode::Jnz, 0);
        assert_eq!(c.ip, 2);

        let mut c = Computer {
            registers: [1, 0, 0],
            ip: 0,
        };
        _ = c.do_instruction(Opcode::Jnz, 0);
        assert_eq!(c.ip, 0);
    }

    #[test]
    fn test_bxc() {
        let mut c = Computer {
            registers: [0, 0b100, 0b11],
            ip: 0,
        };
        _ = c.do_instruction(Opcode::Bxc, 0);
        assert_eq!(c.registers[B], 0b111);
    }

    #[test]
    fn test_bdv() {
        let mut c = Computer {
            registers: [201, 1, 0],
            ip: 0,
        };
        _ = c.do_instruction(Opcode::Bdv, 2);
        assert_eq!(c.registers, [201, 50, 0]);

        let mut c = Computer {
            registers: [201, 1, 0],
            ip: 0,
        };
        _ = c.do_instruction(Opcode::Bdv, 5);
        assert_eq!(c.registers, [201, 100, 0]);
    }

    #[test]
    fn test_cdv() {
        let mut c = Computer {
            registers: [201, 1, 0],
            ip: 0,
        };
        _ = c.do_instruction(Opcode::Cdv, 2);
        assert_eq!(c.registers, [201, 1, 50]);

        let mut c = Computer {
            registers: [201, 1, 0],
            ip: 0,
        };
        _ = c.do_instruction(Opcode::Cdv, 5);
        assert_eq!(c.registers, [201, 1, 100]);
    }
}
