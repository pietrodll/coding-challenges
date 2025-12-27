use super::DayResult;

type Problem<'a> = (Vec<&'a str>, char);

fn parse_input<'a>(data: &'a String) -> Vec<Problem<'a>> {
    let lines: Vec<_> = data.lines().collect();
    let symbols_line = lines[lines.len() - 1];

    let mut previous_symbol = ' ';
    let mut previous_symbol_idx = 0;

    let mut problems: Vec<Problem> = Vec::new();

    for (idx, chr) in symbols_line.char_indices() {
        if !chr.is_whitespace() {
            if !previous_symbol.is_whitespace() {
                let numbers_str: Vec<&str> = lines[0..lines.len() - 1]
                    .iter()
                    .map(|line| &line[previous_symbol_idx..idx - 1])
                    .collect();

                problems.push((numbers_str, previous_symbol));
            }

            previous_symbol = chr;
            previous_symbol_idx = idx;
        }
    }

    // Add the last problem
    let numbers_str: Vec<&str> = lines[0..lines.len() - 1]
        .iter()
        .map(|line| &line[previous_symbol_idx..])
        .collect();
    problems.push((numbers_str, previous_symbol));

    problems
}

fn compute(numbers: &Vec<u64>, symbol: char) -> u64 {
    match symbol {
        '+' => numbers.iter().sum(),

        '*' => numbers
            .iter()
            .fold(1, |product, element| product * *element),
        _ => panic!("Unexpected symbol {}", symbol),
    }
}

fn read_numbers<'a>(numbers_str: &Vec<&'a str>) -> Vec<u64> {
    numbers_str
        .iter()
        .map(|num_str| num_str.trim().parse::<u64>().unwrap())
        .collect()
}

fn compute_total<'a>(problems: &Vec<Problem<'a>>, read_fn: fn(&Vec<&'a str>) -> Vec<u64>) -> u64 {
    problems
        .iter()
        .map(|pb| compute(&read_fn(&pb.0), pb.1))
        .sum()
}

fn first_part(problems: &Vec<Problem>) -> u64 {
    compute_total(problems, read_numbers)
}

fn read_numbers_vertical<'a>(numbers_str: &Vec<&'a str>) -> Vec<u64> {
    let mut numbers: Vec<u64> = Vec::new();

    let chars: Vec<Vec<char>> = numbers_str
        .iter()
        .map(|line| line.chars().collect())
        .collect();

    let mut idx: isize = chars[0].len() as isize - 1;

    while idx >= 0 {
        let number = chars
            .iter()
            .map(|num_chars| num_chars[idx as usize])
            .filter(|num_char| !num_char.is_whitespace())
            .fold(0, |acc, num_char| {
                10 * acc + num_char.to_digit(10).unwrap() as u64
            });

        numbers.push(number);
        idx -= 1;
    }

    numbers
}

fn second_part(problems: &Vec<Problem>) -> u64 {
    compute_total(problems, read_numbers_vertical)
}

pub fn run(data: &String) -> DayResult {
    let problems = parse_input(data);

    DayResult {
        first_part: first_part(&problems).to_string(),
        second_part: second_part(&problems).to_string(),
    }
}

#[cfg(test)]
mod tests {
    use std::vec;

    use super::*;

    const INPUT_DATA: &str = "123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  ";

    #[test]
    fn test_parse_input() {
        let data = String::from(INPUT_DATA);
        let problems = parse_input(&data);
        assert_eq!(problems[0], (vec!["123", " 45", "  6"], '*'));
        assert_eq!(problems[1], (vec!["328", "64 ", "98 "], '+'));
        assert_eq!(problems[2], (vec![" 51", "387", "215"], '*'));
        assert_eq!(problems[3], (vec!["64 ", "23 ", "314"], '+'));
    }

    #[test]
    fn test_first_part() {
        let data = String::from(INPUT_DATA);
        let problems = parse_input(&data);
        assert_eq!(first_part(&problems), 4277556);
    }

    #[test]
    fn test_second_part() {
        let data = String::from(INPUT_DATA);
        let problems = parse_input(&data);
        assert_eq!(second_part(&problems), 3263827);
    }
}
