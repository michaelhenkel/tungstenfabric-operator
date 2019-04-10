cat role.yaml > manifest.yaml
echo "---" >> manifest.yaml
cat service_account.yaml >> manifest.yaml
echo "---" >> manifest.yaml
cat role_binding.yaml >> manifest.yaml
for i in `ls crds/*crd.yaml`
do
	echo "---" >> manifest.yaml
	cat ${i} >> manifest.yaml
done
echo "---" >> manifest.yaml
cat operator.yaml >> manifest.yaml
for i in `ls crds/*cr.yaml`
do
        echo "---" >> manifest.yaml
        cat ${i} >> manifest.yaml
done
