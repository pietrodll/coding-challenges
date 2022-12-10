use std::collections::HashSet;

fn parse_input(data: &String) -> Vec<(&str, &str)> {
    data.lines()
        .map(|line| {
            let half = line.len() / 2;
            (&line[..half], &line[half..])
        })
        .collect()
}

fn find_common_elements(strings: &[&str]) -> HashSet<char> {
    if strings.len() == 0 {
        return HashSet::new();
    }

    let mut common: HashSet<char> = HashSet::from_iter(strings[0].chars());

    for string in &strings[1..] {
        let string_set: HashSet<char> = HashSet::from_iter(string.chars());
        common = common.intersection(&string_set).map(|chr| *chr).collect();
    }

    return common;
}

fn get_priority(chr: char) -> u32 {
    if chr.is_uppercase() {
        (chr as u32) - ('A' as u32) + 27
    } else {
        (chr as u32) - ('a' as u32) + 1
    }
}

pub fn run_first_part(data: &String) -> String {
    parse_input(data)
        .iter()
        .map(
            |tup| match find_common_elements(&[tup.0, tup.1]).iter().next() {
                Some(chr) => get_priority(*chr),
                None => 0,
            },
        )
        .sum::<u32>()
        .to_string()
}

pub fn run_second_part(data: &String) -> String {
    data.lines()
        .collect::<Vec<&str>>()
        .chunks(3)
        .flat_map(find_common_elements)
        .map(get_priority)
        .sum::<u32>()
        .to_string()
}
