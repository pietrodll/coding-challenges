fn parse_input(data: &String) -> Vec<Vec<i32>> {
    return data
        .trim()
        .split("\n\n")
        .map(|str| {
            str.lines()
                .map(|num_str| num_str.parse().unwrap())
                .collect()
        })
        .collect();
}

pub fn run_first_part(data: &String) -> String {
    parse_input(data)
        .iter()
        .map(|sizes| sizes.iter().sum())
        .reduce(|acc: i32, item| if acc > item { acc } else { item })
        .unwrap_or(0)
        .to_string()
}

pub fn run_second_part(data: &String) -> String {
    let mut elf_totals: Vec<i32> = parse_input(data)
        .iter()
        .map(|sizes| sizes.iter().sum())
        .collect();

    elf_totals.sort_unstable();
    elf_totals.reverse();

    let res = elf_totals.get(0).unwrap() + elf_totals.get(1).unwrap() + elf_totals.get(2).unwrap();
    return res.to_string();
}
