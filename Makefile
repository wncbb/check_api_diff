ck:
	go run cmd/checker/main.go  \
			-envFile configs/env.json \
			-reqFile configs/req.json \
			-logLevel info \
			-onlinePOIHost 10.10.53.127:8005 \
			-debugPOIHost 127.0.0.1:7171 \
			-outputDir ./output

ss:
	go run cmd/simpleserver/main.go
	
