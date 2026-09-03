package image

type PartialTemplates struct {
	MaleOld     *PartialHumanOptions `json:"male_old,omitempty"`
	FemaleOld   *PartialHumanOptions `json:"female_old,omitempty"`
	FemaleYoung *PartialHumanOptions `json:"female_young,omitempty"`
	MaleYoung   *PartialHumanOptions `json:"male_young,omitempty"`
	Child       *PartialHumanOptions `json:"child,omitempty"`
}
