struct Range {
    from: i32,
    to: i32,
}

impl Range {
    fn includes(&self, other: &Range) -> bool {
        return self.from <= other.from && self.to >= other.to;
    }
}

fn parse_range(range: &str) -> Range {
    let mut it = range
        .split('-')
        .map(|num_str| num_str.parse::<i32>().unwrap());

    Range {
        from: it.next().unwrap(),
        to: it.next().unwrap(),
    }
}

fn parse_line(line: &str) -> (Range, Range) {
    let mut split = line.split(',');
    let (left, right) = (split.next().unwrap(), split.next().unwrap());

    (parse_range(left), parse_range(right))
}

pub fn run_first_part(data: &String) -> String {
    data.lines()
        .map(parse_line)
        .filter(|(first, second)| first.includes(second) || second.includes(first))
        .count()
        .to_string()
}

pub fn run_second_part(data: &String) -> String {
    data.lines()
        .map(parse_line)
        .filter(|(first, second)| {
            if first.from <= second.from {
                first.to >= second.from
            } else {
                second.to >= first.from
            }
        })
        .count()
        .to_string()
}
