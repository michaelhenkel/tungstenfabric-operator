<<<<<<< HEAD
package ports

=======
>>>>>>> v0.0.4
/*
 Package ports contains the functionality to Listing, Searching, Creating, Updating,
 and Deleting of bare metal Port resources

 API reference: https://developer.openstack.org/api-ref/baremetal/#ports-ports


<<<<<<< HEAD
 	// Example to List Ports with Detail
 	ports.ListDetail(client, ports.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
=======
Example to List Ports with Detail

 	ports.ListDetail(client, nil).EachPage(func(page pagination.Page) (bool, error) {
>>>>>>> v0.0.4
 		portList, err := ports.ExtractPorts(page)
 		if err != nil {
 			return false, err
 		}

 		for _, n := range portList {
 			// Do something
 		}

 		return true, nil
 	})

<<<<<<< HEAD
 	// Example to List Ports
 	ports.List(client, ports.ListOpts{
 		Limit: 10,
 	}).EachPage(func(page pagination.Page) (bool, error) {
=======
Example to List Ports

	listOpts := ports.ListOpts{
 		Limit: 10,
	}

 	ports.List(client, listOpts).EachPage(func(page pagination.Page) (bool, error) {
>>>>>>> v0.0.4
 		portList, err := ports.ExtractPorts(page)
 		if err != nil {
 			return false, err
 		}

 		for _, n := range portList {
 			// Do something
 		}

 		return true, nil
 	})

<<<<<<< HEAD
 	// Example to Create a Port
 	createPort, err := ports.Create(client, ports.CreateOpts{
 		NodeUUID: "e8920409-e07e-41bb-8cc1-72acb103e2dd",
		Address: "00:1B:63:84:45:E6",
    PhysicalNetwork: "my-network",
 	}).Extract()
=======
Example to Create a Port

	createOpts := ports.CreateOpts{
 		NodeUUID: "e8920409-e07e-41bb-8cc1-72acb103e2dd",
		Address: "00:1B:63:84:45:E6",
    PhysicalNetwork: "my-network",
	}

 	createPort, err := ports.Create(client, createOpts).Extract()
>>>>>>> v0.0.4
 	if err != nil {
 		panic(err)
 	}

<<<<<<< HEAD
 	// Example to Get a Port
=======
Example to Get a Port

>>>>>>> v0.0.4
 	showPort, err := ports.Get(client, "c9afd385-5d89-4ecb-9e1c-68194da6b474").Extract()
 	if err != nil {
 		panic(err)
 	}

<<<<<<< HEAD
 	// Example to Update a Port
 	updatePort, err := ports.Update(client, "c9afd385-5d89-4ecb-9e1c-68194da6b474", ports.UpdateOpts{
=======
Example to Update a Port

	updateOpts := ports.UpdateOpts{
>>>>>>> v0.0.4
 		ports.UpdateOperation{
 			Op:    ReplaceOp,
 			Path:  "/address",
 			Value: "22:22:22:22:22:22",
 		},
<<<<<<< HEAD
 	}).Extract()
=======
	}

 	updatePort, err := ports.Update(client, "c9afd385-5d89-4ecb-9e1c-68194da6b474", updateOpts).Extract()
>>>>>>> v0.0.4
 	if err != nil {
 		panic(err)
 	}

<<<<<<< HEAD
 	// Example to Delete a Port
 	err = ports.Delete(client, "c9afd385-5d89-4ecb-9e1c-68194da6b474").ExtractErr()
 	if err != nil {
 		panic(err)

*/
=======
Example to Delete a Port

 	err = ports.Delete(client, "c9afd385-5d89-4ecb-9e1c-68194da6b474").ExtractErr()
 	if err != nil {
 		panic(err)
	}

*/
package ports
>>>>>>> v0.0.4
