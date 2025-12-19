pub struct Matrix<T> {
    data: Vec<T>,
    rows: usize,
    cols: usize,
}

// Use isize to handle negative offsets
const DELTAS: [(isize, isize); 8] = [
    (-1, -1), // top-left
    (-1, 0),  // top
    (-1, 1),  // top-right
    (0, -1),  // left
    (0, 1),   // right
    (1, -1),  // bottom-left
    (1, 0),   // bottom
    (1, 1),   // bottom-right
];

impl<T> Matrix<T> {
    pub fn new(rows: usize, cols: usize, default: T) -> Self
    where
        T: Clone,
    {
        Self {
            data: vec![default; rows * cols],
            rows,
            cols,
        }
    }

    pub fn rows(&self) -> usize {
        self.rows
    }

    pub fn cols(&self) -> usize {
        self.cols
    }

    pub fn get(&self, row: usize, col: usize) -> Option<&T> {
        if row < self.rows && col < self.cols {
            Some(&self.data[row * self.cols + col])
        } else {
            None
        }
    }

    pub fn set(&mut self, row: usize, col: usize, value: T) {
        if row < self.rows && col < self.cols {
            self.data[row * self.cols + col] = value;
        } else {
            panic!(
                "out of bounds: position ({}, {}) for matrix of size ({}, {})",
                row, col, self.rows, self.cols
            )
        }
    }

    pub fn neighbors(&self, row: usize, col: usize) -> impl Iterator<Item = &T> {
        let rows = self.rows as isize;
        let cols = self.cols as isize;

        DELTAS.iter().filter_map(move |(dr, dc)| {
            let new_row = row as isize + dr;
            let new_col = col as isize + dc;

            if new_row >= 0 && new_row < rows && new_col >= 0 && new_col < cols {
                self.get(new_row as usize, new_col as usize)
            } else {
                None
            }
        })
    }

    pub fn entries_iter(&self) -> impl Iterator<Item = ((usize, usize), &T)> {
        (0..self.rows).flat_map(move |row| {
            (0..self.cols).map(move |col| ((row, col), self.get(row, col).unwrap()))
        })
    }
}

impl<T> FromIterator<Vec<T>> for Matrix<T> {
    fn from_iter<I: IntoIterator<Item = Vec<T>>>(iter: I) -> Self {
        let mut data: Vec<T> = Vec::new();
        let mut cols: usize = 0;
        let mut rows: usize = 0;
        let mut rows_iterator = iter.into_iter();

        while let Some(mut row) = rows_iterator.next() {
            if cols == 0 {
                cols = row.len();
            } else if row.len() != cols {
                panic!("Inconsistent length of rows")
            }

            rows += 1;
            data.append(&mut row);
        }

        Self {
            data: data,
            rows: rows,
            cols: cols,
        }
    }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_neighbors() {
        let data: Vec<Vec<i32>> = vec![vec![1, 2, 3], vec![4, 5, 6], vec![7, 8, 9]];
        let matrix: Matrix<i32> = data.into_iter().collect();

        assert_eq!(
            matrix.neighbors(1, 1).map(|num| *num).collect::<Vec<_>>(),
            vec![1, 2, 3, 4, 6, 7, 8, 9]
        )
    }

    #[test]
    fn test_collect() {
        let data: Vec<Vec<i32>> = vec![vec![1, 2, 3], vec![4, 5, 6]];
        let matrix: Matrix<i32> = data.into_iter().collect();

        assert_eq!(matrix.rows, 2);
        assert_eq!(matrix.cols, 3);
        assert_eq!(matrix.data, vec![1, 2, 3, 4, 5, 6]);
    }
}
