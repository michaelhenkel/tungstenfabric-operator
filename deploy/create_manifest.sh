cat role.yaml > manifest.yaml
echo "---" >> manifest.yaml
cat service_account.yaml >> manifest.yaml
echo "---" >> manifest.yaml
cat role_binding.yaml >> manifest.yaml
echo "---" >> manifest.yaml
cat crds/tf_v1alpha1_tungstenfabricconfig_crd.yaml >> manifest.yaml
echo "---" >> manifest.yaml
cat operator.yaml >> manifest.yaml
echo "---" >> manifest.yaml
cat crds/tf_v1alpha1_tungstenfabricconfig_cr.yaml >> manifest.yaml
for i in `ls crds/*cr.yaml |grep -v tf_v1alpha1_tungstenfabricconfig_cr.yaml`
do
        echo "---" >> manifest.yaml
        cat ${i} >> manifest.yaml
done
