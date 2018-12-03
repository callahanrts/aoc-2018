include Day1_data;

let b = Array.fold_left((+), 0, changes);
Printf.printf("Day 1: %d", b)
