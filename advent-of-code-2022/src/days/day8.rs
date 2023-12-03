fn parse_input(data: &String) -> Vec<Vec<u32>> {
    data.trim()
        .lines()
        .map(|line| {
            line.chars()
                .map(|char| char.to_digit(10).unwrap())
                .collect()
        })
        .collect()
}

fn find_visible(vect: &Vec<u32>, is_visible: &mut Vec<bool>) {
    let mut max_height: u32 = 0;

    for i in 0..vect.len() {
        let tree_height = vect.get(i).unwrap();

        if *tree_height > max_height {
            *(is_visible.get_mut(i).unwrap()) = true;
            max_height = *tree_height;
        }
    }

    max_height = 0;

    for i in vect.len()..0 {
        let tree_height = vect.get(i).unwrap();

        if *tree_height > max_height {
            *(is_visible.get_mut(i).unwrap()) = true;
            max_height = *tree_height;
        }
    }
}

fn count_visible(grid: &Vec<Vec<u32>>) -> u32 {
    let num_rows = grid.len();
    let num_cols = grid.get(0).unwrap().len();

    let mut is_visible: Vec<Vec<bool>> = Vec::new();

    for row in grid {
        is_visible.push(vec![false; row.len()]);
    }

    for i in 0..num_rows {
        find_visible(grid.get(i).unwrap(), is_visible.get_mut(i).unwrap());
    }

    for j in 0..num_cols {
        find_visible(grid.get(index), is_visible)
    }

    is_visible
        .into_iter()
        .flatten()
        .filter(|v| *v)
        .count()
        .try_into()
        .unwrap()
}

pub fn run_first_part(data: &String) -> String {
    count_visible(&parse_input(data)).to_string()
}

pub fn run_second_part(data: &String) -> String {
    String::new()
}
