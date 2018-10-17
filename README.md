# fpga

## aws f1 instance
https://github.com/aws/aws-fpga/blob/master/hdk/README.md


## load bitstream to local machine
open vivado
open hardware manager
autoconnect the machine
program device
choose bitstream in ~/Desktop/fpga
restart the computer
load driver from ~/Desktop/pcie_driver/Xilinx65444/tests/load_driver.sh
run the host code in ~/Desktop/fpga/lrcn
