package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/joyent/solidfire-sdk/api"
)

// Example for setting a middleware for recording raw requests and responses
func middlewareExample(c *api.Client) (err error) {
	c.HTTPClient.OnAfterResponse(func(client *resty.Client, resp *resty.Response) error {
		req := *resp.Request
		// N.B. - Using fmt.Printf here for simplicity and to avoid a dependency but a real
		// client would likely use a logger.
		fmt.Printf("Request: %#v\n", req)
		fmt.Printf("Response Header: %#v\n", resp.Header())
		fmt.Printf("Response Body: %#v\n", string(resp.Body()))
		return nil
	})
	return nil
}

func volumeExamples(c *api.Client, accountId int64) (volume *api.Volume, err error) {
	ctx := context.Background()
	request := api.CreateVolumeRequest{
		Name:       "solidfire-sdk-example",
		AccountID:  accountId,
		TotalSize:  1000000000,
		Enable512e: true,
	}
	createdVolume, err := c.CreateVolume(ctx, request)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Created volume with ID %d\n", createdVolume.VolumeID)
	volumeId := createdVolume.VolumeID
	req := api.ModifyVolumeRequest{
		VolumeID:  volumeId,
		TotalSize: 1000001000,
	}
	_, err = c.ModifyVolume(ctx, req)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Modified volume %s\n", createdVolume.Name)
	volume, err = c.GetVolumeById(ctx, volumeId)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Found created volume via ListVolumes call %s\n", volume.Name)
	return volume, nil
}

func snapshotExamples(c *api.Client, volumeId int64) (snapshots []api.Snapshot, err error) {
	ctx := context.Background()
	csr := api.CreateSnapshotRequest{
		VolumeID: 3576,
	}
	createdSnapshot, err := c.CreateSnapshot(ctx, csr)
	if err != nil {
		return nil, err
	}
	snapshotId := createdSnapshot.SnapshotID
	fmt.Printf("Created snapshot %d for volume %d\n", snapshotId, volumeId)
	msr := api.ModifySnapshotRequest{
		SnapshotID: snapshotId,
		Name:       "solidfire-sdk-example-snapshot",
	}
	_, err = c.ModifySnapshot(ctx, msr)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Modified snapshot %d\n", snapshotId)
	snapshots, err = c.GetSnapshotsByVolumeId(ctx, volumeId)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Found created snapshot %d\n", snapshotId)
	err = c.DeleteSnapshot(ctx, snapshotId)
	fmt.Printf("Deleted snapshot %d\n", snapshotId)
	return snapshots, err
}

func main() {
	host := os.Getenv("SOLIDFIRE_HOST")
	username := os.Getenv("SOLIDFIRE_USER")
	password := os.Getenv("SOLIDFIRE_PASS")
	if host == "" || username == "" || password == "" {
		fmt.Println("Environment variables SOLIDFIRE_HOST, SOLIDFIRE_USER, and SOLIDFIRE_PASS must be set")
		return
	}

	c, err := api.BuildClient(host, username, password, "12.3", 443, 3)
	if err != nil {
		fmt.Printf("Error connecting: %s\n", err)
		panic(err)
	}
	err = middlewareExample(c)
	if err != nil {
		fmt.Printf("Error with setting resty middleware: %s\n", err)
		panic(err)
	}

	ctx := context.Background()
	lar := api.ListAccountsRequest{}
	accounts, err := c.ListAccounts(ctx, lar)
	if err != nil || len(accounts) == 0 {
		fmt.Printf("Error listing accounts: %s\n", err)
		panic(err)
	}
	fmt.Printf("Found %d accounts\n", len(accounts))
	accountId := accounts[0].AccountID

	volume, err := volumeExamples(c, accountId)
	if err != nil {
		fmt.Printf("Error with volume examples: %s\n", err)
		panic(err)
	}
	_, err = snapshotExamples(c, volume.VolumeID)
	if err != nil {
		fmt.Printf("Error with snapshot examples: %s\n", err)
		panic(err)
	}
	_, err = c.DeleteVolume(ctx, volume.VolumeID)
	fmt.Printf("Deleted created volume %s\n", volume.Name)
	if err != nil {
		fmt.Print(err)
	}
}
