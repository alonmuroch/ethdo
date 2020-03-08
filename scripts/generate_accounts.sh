#!/bin/bash
# Tested using bash version 4.1.5
for ((i=1;i<=1000;i++));
do

   ethdo account create --account=Validators/$i --passphrase=12345

   x=$(ethdo validator depositdata \
                   --validatoraccount=Validators/$i \
                   --withdrawalaccount=Withdrawal/Primary \
                   --depositvalue=3.2Ether \
                   --passphrase=12345)

   ethereal beacon deposit \
      --network=goerli \
      --data=$x \
      --from=0xeFc0D9A1C9c75837F842abc90690B7c325717D59 \
      --privatekey=$PRIVATEKEY

   echo "Sent deposit for validator $i"
done