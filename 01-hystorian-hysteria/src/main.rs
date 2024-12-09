use std::fs;
use std::io::{self, BufRead};
use regex::Regex;
use std::collections::HashMap;

fn read_file(filename: &str) -> io::Result<Vec<String>> {
    let file = fs::File::open(filename)?;
    let reader = io::BufReader::new(file);
    return reader.lines().collect();
}

fn part1() -> Result<(), Box<dyn std::error::Error>> {
    let lines = read_file("./src/input.txt")?;

    let re = Regex::new(r"[0-9]+").unwrap();

    let mut arr1: Vec<i32> = Vec::new();
    let mut arr2: Vec<i32> = Vec::new();
    for line in lines {
        println!("Matching line {}", line);
        let mut index = 0;
        for mat in re.find_iter(&line) {
            let n = mat.as_str().parse::<i32>().unwrap();
            if index == 0 {
                arr1.push(n);
            } else if index == 1 {
                arr2.push(n);
            } else {
                println!("Something is wrong");
            }

            index = index + 1;
        }
    }

    arr1.sort_by(|a, b| a.cmp(b));
    arr2.sort_by(|a, b| a.cmp(b));

    let mut result = 0;

    for i in 0..arr1.len() {
        let a = arr1[i];
        let b = arr2[i];

        result = result + (b - a).abs();
    }

    println!("Result {}", result);
    return Ok(());
}

fn part2() -> Result<(), Box<dyn std::error::Error>> {
    let lines = read_file("./src/input.txt")?;

    let re = Regex::new(r"[0-9]+").unwrap();

    let mut arr1: Vec<i32> = Vec::new();
    let mut arr2: Vec<i32> = Vec::new();
    for line in lines {
        let mut index = 0;
        for mat in re.find_iter(&line) {
            let n = mat.as_str().parse::<i32>().unwrap();
            if index == 0 {
                arr1.push(n);
            } else if index == 1 {
                arr2.push(n);
            } else {
                println!("Something is wrong");
            }

            index = index + 1;
        }
    }

    let mut freq_map: HashMap<i32, i32> = HashMap::new();

    for n in arr2 {
        if let Some(value) = freq_map.get(&n) {
            freq_map.insert(n, value + 1);
        } else {
            freq_map.insert(n, 1);
        }
    }

    let mut result = 0;
    for n in arr1 {
        if let Some(freq) = freq_map.get(&n) {
            result = result + (n * freq);
        }
    }

    println!("{}", result);
    return Ok(());
}

fn main() -> Result<(), Box<dyn std::error::Error>> {
    part2()?;
    return Ok(());
}
