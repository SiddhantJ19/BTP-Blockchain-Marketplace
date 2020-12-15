
# app_cmd_dir="/run/media/coderdude/Adwait/Projects/btp_fabric/tutorials/fabric-samples/test-network/btp_cmds/app_cmds/"
app_cmd_dir="/home/siddhant/HLFabric/fabric-samples/Blockchain-marketplace/Application/app_script_cmds/app_cmds"
# echo "Enrolling org1"
source $app_cmd_dir/org1_enroll_and_connect.sh


# testEvent
# source $app_cmd_dir/testEvent.sh

# echo "Enrolling org2"
# source $app_cmd_dir/org2_enroll_and_connect.sh

# # echo "Connection gateway from org1"
# # source $app_cmd_dir/org1_connect.sh

# echo "\nCreating org1 devices"
# source $app_cmd_dir/org1_create_devices.sh

# echo "\nCreating org2 devices"
# source $app_cmd_dir/org2_create_devices.sh

# # share from org1 - org2
# source $app_cmd_dir/org1_add_data.sh
# source $app_cmd_dir/org1_add_data.sh

# source $app_cmd_dir/sharing_org1_to_2.sh
# source $app_cmd_dir/org1_add_data.sh
# source $app_cmd_dir/org1_add_data.sh
# echo "\n\n adding data after revoking access"
source $app_cmd_dir/org1_add_data.sh


# # share from org2 - org1

# source $app_cmd_dir/org2_add_data.sh
# source $app_cmd_dir/org2_add_data.sh

# source $app_cmd_dir/sharing_org2_to_1.sh

# source $app_cmd_dir/org2_add_data.sh


# # org1 get shared data

# source $app_cmd_dir/org1_get_shared_data.sh

# # org2 get shared data
# source $app_cmd_dir/org2_get_shared_data.sh
