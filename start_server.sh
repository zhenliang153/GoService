
pkill -9 -f Service
sleep 2
./Service 1>err 2>&1 &
