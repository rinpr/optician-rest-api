package eyedata

type SpectaclePrescription struct {
	RightEye Eye `json:"right_eye,omitempty" bson:"right_eye,omitempty"`
	LeftEye  Eye `json:"left_eye,omitempty" bson:"left_eye,omitempty"`
}

type Eye struct {
	Sphere   float32 `json:"sphere,omitempty" bson:"sphere,omitempty"`
	Cylinder float32 `json:"cylinder,omitempty" bson:"cylinder,omitempty"`
	Axis     float32 `json:"axis,omitempty" bson:"axis,omitempty"`
	Prism    float32 `json:"prism,omitempty" bson:"prism,omitempty"`
	Add      float32 `json:"add,omitempty" bson:"add,omitempty"`
	MonoVa   string  `json:"mono_va,omitempty" bson:"mono_va,omitempty"`
	DPD      int     `json:"dpd,omitempty" bson:"dpd,omitempty"`
	NPD      int     `json:"npd,omitempty" bson:"npd,omitempty"`
	FH       int     `json:"fh,omitempty" bson:"fh,omitempty"`
}
