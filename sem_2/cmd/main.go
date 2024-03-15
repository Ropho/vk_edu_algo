package main

import "vk_algo/sem_2"

func main() {

	head := sem_2.ConstructList([]sem_2.Note{
		{
			Val:       11,
			NextIndex: 1,
		},
		{
			Val:       3,
			NextIndex: 2,
		},
		{
			Val:       8,
			NextIndex: 3,
		},
		{
			Val:       9,
			NextIndex: 4,
		},
		{
			Val:       6,
			NextIndex: 5,
		},
		{
			Val:       11,
			NextIndex: 6,
		},
		{
			Val:       16,
			NextIndex: 7,
		},
		{
			Val:       17,
			NextIndex: 2,
		},
	})

	sem_2.PrintAll(head)
}
