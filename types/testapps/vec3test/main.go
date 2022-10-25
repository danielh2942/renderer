package main

import (
	"encoding/json"
	"fmt"

	"github.com/danielh2942/renderer/types"
)

// Verify that vector3d can be encoded as json 

func main() {
	mVec3d := types.Vector3d{
		Z: 3.0,
	}

	mVec3d.Vector2d = types.Vector2d{
		X: 1.0,
		Y: 2.0,
	}

	fmt.Println(mVec3d)

	outp, _ := json.Marshal(mVec3d)

	fmt.Println(string(outp))

	var outVec types.Vector3d

	json.Unmarshal(outp,&outVec)

	fmt.Println(outVec)
}
