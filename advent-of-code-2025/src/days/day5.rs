use super::DayResult;

type Range = (u64, u64);

fn parse_input(data: &String) -> (Vec<Range>, Vec<u64>) {
    let split: Vec<&str> = data.trim().splitn(2, "\n\n").collect();

    let ranges: Vec<(u64, u64)> = split[0]
        .lines()
        .map(|line| {
            let parsed: Vec<u64> = line
                .splitn(2, '-')
                .map(|s| s.parse::<u64>().unwrap())
                .collect();
            (parsed[0], parsed[1])
        })
        .collect();

    let ingredients: Vec<u64> = split[1].lines().map(|line| line.parse().unwrap()).collect();

    (ranges, ingredients)
}

fn first_part(ranges: &Vec<Range>, ingredients: &Vec<u64>) -> u32 {
    ingredients
        .iter()
        .filter(|ingredient| {
            ranges
                .iter()
                .any(|range| range.0 <= **ingredient && **ingredient <= range.1)
        })
        .count() as u32
}

/// Returns an index i such that `ranges[i]` contains the start of `range`. If
/// no range overlaps, returns the index where `range` should be inserted,
/// assuming that other elements as shifted to the right.
fn find_overlapping_range_index(ranges: &Vec<Range>, range: &Range) -> usize {
    let mut l: usize = 0;
    let mut r: usize = ranges.len();

    while l < r {
        let m = (l + r) / 2;
        let m_range = ranges[m];

        if m_range.0 <= range.0 && range.0 <= m_range.1 {
            return m;
        } else if range.0 < m_range.0 {
            // the range is lower than the middle one, move to the left
            r = m;
        } else {
            l = m + 1;
        }
    }

    return l;
}

/// Inserts `range` in `ranges` such that `ranges` keeps being a list of disjoint
/// ranges sorted in ascending order
fn insert_merge(ranges: &mut Vec<Range>, range: Range) {
    let overlapping_idx = find_overlapping_range_index(ranges, &range);

    if overlapping_idx == ranges.len() {
        ranges.push(range);
    } else {
        let found_range = ranges[overlapping_idx];

        if range.1 < found_range.0 {
            // If the end of `range` is before `found_range`, it means they are disjoint:
            // insert the range at the found index
            ranges.insert(overlapping_idx, range);
        } else {
            // `range` and `found_range` are not disjoint: set their union as the element at index `overlapping_idx`
            ranges[overlapping_idx] = (
                std::cmp::min(found_range.0, range.0),
                std::cmp::max(found_range.1, range.1),
            );

            // we might have extended the end, so the range might overlap with the next one and need to be merged
            while overlapping_idx < ranges.len() - 1
                && ranges[overlapping_idx].1 >= ranges[overlapping_idx + 1].0
            {
                let merged = (ranges[overlapping_idx].0, ranges[overlapping_idx + 1].1);
                ranges.splice(overlapping_idx..=overlapping_idx + 1, [merged]);
            }
        }
    }
}

fn second_part(ranges: &Vec<Range>) -> u64 {
    let mut fresh_ranges: Vec<Range> = Vec::new();

    for range in ranges {
        insert_merge(&mut fresh_ranges, range.clone());
    }

    fresh_ranges.iter().map(|range| range.1 - range.0 + 1).sum()
}

pub fn run(data: &String) -> DayResult {
    let (ranges, ingredients) = parse_input(data);

    DayResult {
        first_part: first_part(&ranges, &ingredients).to_string(),
        second_part: second_part(&ranges).to_string(),
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT_DATA: &str = "3-5
10-14
16-20
12-18

1
5
8
11
17
32";

    #[test]
    fn test_first_part() {
        let (ranges, ingredients) = parse_input(&String::from(INPUT_DATA));
        assert_eq!(first_part(&ranges, &ingredients), 3)
    }

    #[test]
    fn test_find_overlapping_range_index() {
        let ranges: Vec<Range> = vec![(2, 4), (7, 9), (13, 20)];

        assert_eq!(find_overlapping_range_index(&ranges, &(1, 5)), 0);
        assert_eq!(find_overlapping_range_index(&ranges, &(6, 10)), 1);
        assert_eq!(find_overlapping_range_index(&ranges, &(3, 4)), 0);
        assert_eq!(find_overlapping_range_index(&ranges, &(5, 6)), 1);
        assert_eq!(find_overlapping_range_index(&ranges, &(0, 1)), 0);
        assert_eq!(find_overlapping_range_index(&ranges, &(25, 30)), 3);
    }

    #[test]
    fn test_second_part() {
        let (ranges, _) = parse_input(&String::from(INPUT_DATA));
        assert_eq!(second_part(&ranges), 14);
    }
}
