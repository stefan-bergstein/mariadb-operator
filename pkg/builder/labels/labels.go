package builder

const (
	appLabel       = "app.kubernetes.io/name"
	instanceLabel  = "app.kubernetes.io/instance"
	componentLabel = "app.kubernetes.io/component"
	releaseLabel   = "release"
)

type LabelsBuilder struct {
	labels map[string]string
}

func NewLabelsBuilder() *LabelsBuilder {
	return &LabelsBuilder{
		labels: map[string]string{},
	}
}

func (b *LabelsBuilder) WithApp(app string) *LabelsBuilder {
	b.labels[appLabel] = app
	return b
}

func (b *LabelsBuilder) WithInstance(instance string) *LabelsBuilder {
	b.labels[instanceLabel] = instance
	return b
}

func (b *LabelsBuilder) WithComponent(component string) *LabelsBuilder {
	b.labels[componentLabel] = component
	return b
}

func (b *LabelsBuilder) WithRelease(release string) *LabelsBuilder {
	b.labels[releaseLabel] = release
	return b
}

func (b *LabelsBuilder) Build() map[string]string {
	return b.labels
}
