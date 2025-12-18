use super::DayResult;

fn parse_line(line: &str) -> Option<i64> {
    let mut chars = line.chars();

    match chars.next() {
        Some('R') => chars.collect::<String>().parse::<i64>().ok(),
        Some('L') => chars
            .collect::<String>()
            .parse::<i64>()
            .ok()
            // Multiply by -1 when the rotation is to the left
            .map(|val| -1 * val),
        _ => None,
    }
}

fn parse_input(data: &String) -> Vec<i64> {
    data.trim().lines().filter_map(parse_line).collect()
}

fn first_part(values: &Vec<i64>) -> i64 {
    values
        .iter()
        .fold((50, 0), |(sum, zeroes), value| {
            let next_sum = (sum + value).rem_euclid(100);
            (next_sum, zeroes + if next_sum == 0 { 1 } else { 0 })
        })
        .1
}

fn second_part(values: &Vec<i64>) -> i64 {
    values
        .iter()
        .fold((50, 0), |(sum, zeroes), value| {
            let reduced_value = value % 100;
            let turns = value.abs() / 100;
            let next_without_modulo = sum + reduced_value;

            let next_zeroes = zeroes
                + turns
                + if sum != 0
                    && ((reduced_value < 0 && next_without_modulo <= 0)
                        || (reduced_value > 0 && next_without_modulo >= 100))
                {
                    1
                } else {
                    0
                };

            let next_sum = next_without_modulo.rem_euclid(100);

            (next_sum, next_zeroes)
        })
        .1
}

pub fn run(data: &String) -> DayResult {
    let values = parse_input(data);

    let first_part = first_part(&values).to_string();

    let second_part = second_part(&values).to_string();

    DayResult {
        first_part,
        second_part,
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT_DATA: &str = "L68
L30
R48
L5
R60
L55
L1
L99
R14
L82";

    #[test]
    fn test_parse_input() {
        let input = String::from(INPUT_DATA);
        let expected: Vec<i64> = vec![-68, -30, 48, -5, 60, -55, -1, -99, 14, -82];
        assert_eq!(parse_input(&input), expected);
    }

    #[test]
    fn test_first_part() {
        let values = parse_input(&String::from(INPUT_DATA));
        assert_eq!(first_part(&values), 3);
    }

    #[test]
    fn test_second_part() {
        let values = parse_input(&String::from(INPUT_DATA));
        assert_eq!(second_part(&values), 6);

        let value_with_multiple_turns: Vec<i64> =
            vec![-68, -30, 1048, -5, 60, -55, -1, -99, 14, -82];
        assert_eq!(second_part(&value_with_multiple_turns), 16);
    }
}
