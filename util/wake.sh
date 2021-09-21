wake=$(wget -O - --user=<USERNAME> --password=<PASSWORD> <HOME_SERVICE_URL>/v1/wake | jq .wake_device)
if $wake; then
    wakeonlan -i <TARGET_IP> -p 7 <TARGET_MAC>
fi