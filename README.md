


1.在master
cd /home/ubuntu/hlf-kafka
export PATH=${PWD}/bin:${PWD}:$PATH
export FABRIC_CFG_PATH=${PWD}

2.在master
./bin/cryptogen generate --config=./crypto-config.yaml

3.在master
./bin/configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ./network-config/orderer.block

4.在master
./bin/configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./network-config/channel.tx -channelID mychannel

5.在master
./bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./network-config/Org1MSPanchors.tx -channelID mychannel -asOrg Org1MSP

./mac_bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./network-config/Org2MSPanchors.tx -channelID mychannel -asOrg Org2MSP

6.在master
ls crypto-config/peerOrganizations/org1.example.com/ca #获取CA文件名 
vim deployment/docker-compose-kafka.yml  #修改CA文件名，FABRIC_CA_SERVER_CA_KEYFILE

7.在master
docker-compose -f deployment/docker-compose-kafka.yml up -d
#传递文件夹crypto-config 和network-config 到peer0, peer1

8.在peer0: 10.10.70.137
docker-compose -f deployment/docker-compose-peer0.yml up -d
docker-compose -f deployment/docker-compose-cli0.yml up -d

9.在peer1: 10.10.121.27
docker-compose -f deployment/docker-compose-peer1.yml up -d
docker-compose -f deployment/docker-compose-cli1.yml up -d

10.在peer0: 10.10.70.137
docker exec cli peer channel create -o orderer0.example.com:7050 -c mychannel -f /var/hyperledger/configs/channel.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer0.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

docker-compose -f deployment/docker-compose-cli0.yml exec cli peer channel join -b mychannel.block

docker-compose -f deployment/docker-compose-cli0.yml exec cli peer channel update -o orderer0.example.com:7050 -c mychannel -f /var/hyperledger/configs/Org1MSPanchors.tx  --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer0.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

11.在peer0: 10.10.70.137
docker cp cli:/opt/gopath/src/github.com/hyperledger/fabric/peer/mychannel.block .
#传递文件 scp -r mychannel.block ubuntu@10.10.121.27:/home/ubuntu/deployment

12.在peer1: 10.10.121.27
docker cp mychannel.block cli:/opt/gopath/src/github.com/hyperledger/fabric/peer/mychannel.block
rm mychannel.block

docker-compose -f deployment/docker-compose-cli1.yml exec cli peer channel join -b mychannel.block

docker-compose -f deployment/docker-compose-cli1.yml exec cli peer channel update -o orderer0.example.com:7050 -c mychannel -f /var/hyperledger/configs/Org2MSPanchors.tx  --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer0.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

13.在peer0: 10.10.70.137
docker exec -it cli peer chaincode install -n marble -p github.com/chaincode -v 1.0

14.在peer1: 10.10.121.27
docker exec -it cli peer chaincode install -n marble -p github.com/chaincode -v 1.0

15.在peer0: 10.10.70.137
docker exec -it cli peer chaincode instantiate -o orderer0.example.com:7050 -C mychannel -n marble github.com/chaincode -v 1.0 -c '{"Args":["init"]}' --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer0.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

16.
docker exec -it cli peer chaincode invoke -o orderer0.example.com:7050 -n marble -c '{"Args":["initStringSha", "2040"]}' -C mychannel --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer0.example.com/msp/tlscacerts/tlsca.example.com-cert.pem


curl http://10.10.30.136:8080/api/channels/info
curl http://106.75.67.41:8080/api/channels/info



docker exec -it cli peer chaincode instantiate -o orderer0.example.com:7050 -C mychannel -n mycc github.com/chaincode -v v0 -c '{"Args": ["a", "100"]}'



docker exec -it cli peer chaincode invoke -o orderer0.example.com:7050 -n mycc -c '{"Args":["set", "a", "20"]}' -C mychannel

docker-compose -f deployment/docker-compose-peer1.yml down -v
docker-compose -f deployment/docker-compose-cli1.yml down -v

docker-compose -f deployment/docker-compose-peer0.yml down -v --rm all
docker-compose -f deployment/docker-compose-cli0.yml down -v --rm all

docker-compose -f deployment/docker-compose-kafka.yml down -v

export GOROOT=/usr/local/go
export GOPATH=$HOME/Projects/golang
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
