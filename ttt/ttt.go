package ttt

func Play(options *Options) {
    grid := init_grid(options.Token1, options.Token2, options.Empty, options.Dimension)
	print_grid(grid)
}
