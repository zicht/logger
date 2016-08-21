

	
	logger := logger.NewLogger("main", 
	    handlers.NewWriterHandler("main_handler", os.Stdout, logger.INFO)                             // will write all channels 
	    handlers.NewWriterHandler("main_handler", file01, logger.INFO, logger.ChannelName("redis"))   // will write only to redis
	    handlers.NewWriterHandler("main_handler", file02, logger.INFO, logger.ChannelName("mysql"))   // will write only to mysql
	)
	logger.Register("redis")
	logger.Register("mysql")
