use std::fs::File;
use std::io::BufReader;
use std::io::prelude::*;
use aoc_2019_01::find_fuel_requirement;
use aoc_2019_01::find_fuel_fuel_requirement;

fn main() -> std::io::Result<()> {
    let file = File::open("input")?;
    let buf_reader = BufReader::new(file);
    let mut sum_a: i32 = 0;
    let mut sum_b: i32 = 0;
    for line in buf_reader.lines() {
        let i: i32 = line?.parse().unwrap_or(0);
        let base_fuel: i32 = find_fuel_requirement(i);
        sum_a = sum_a + base_fuel;
        sum_b = sum_b + find_fuel_fuel_requirement(base_fuel);
    }
    println!("A: {}", sum_a);
    println!("B: {}", sum_b);
    Ok(())
}
