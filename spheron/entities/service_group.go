package entities

type Group struct {
	Name     string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name" yaml:"name"`
	Services Services `protobuf:"bytes,2,rep,name=services,proto3,castrepeated=Services" json:"services" yaml:"services"`
}

// GetName returns the name of group
func (g Group) GetName() string {
	return g.Name
}

func (g Group) GetResourceUnits() ResourceUnits {
	groups := make(map[uint32]*ResourceUnit)

	for _, svc := range g.Services {
		if _, exists := groups[svc.Resources.ID]; !exists {
			groups[svc.Resources.ID] = &ResourceUnit{
				Resources: svc.Resources,
				Count:     svc.Count,
			}
		} else {
			groups[svc.Resources.ID].Count += svc.Count
		}
	}

	units := make(ResourceUnits, 0, len(groups))

	for i := range groups {
		units = append(units, *groups[i])
	}

	return units
}

type Services []Service

type Service struct {
	Name        string                   `protobuf:"bytes,1,opt,name=name,proto3" json:"name" yaml:"name"`
	Image       string                   `protobuf:"bytes,2,opt,name=image,proto3" json:"image" yaml:"image"`
	Command     []string                 `protobuf:"bytes,3,rep,name=command,proto3" json:"command" yaml:"command"`
	Args        []string                 `protobuf:"bytes,4,rep,name=args,proto3" json:"args" yaml:"args"`
	Env         []string                 `protobuf:"bytes,5,rep,name=env,proto3" json:"env" yaml:"env"`
	Resources   Resources                `protobuf:"bytes,6,opt,name=resources,proto3" json:"resources" yaml:"resources"`
	Count       uint32                   `protobuf:"varint,7,opt,name=count,proto3" json:"count" yaml:"count"`
	Expose      ServiceExposes           `protobuf:"bytes,8,rep,name=expose,proto3,castrepeated=ServiceExposes" json:"expose" yaml:"expose"`
	Params      *ServiceParams           `protobuf:"bytes,9,opt,name=params,proto3" json:"params" yaml:"params"`
	Credentials *ServiceImageCredentials `protobuf:"bytes,10,opt,name=credentials,proto3" json:"credentials" yaml:"credentials"`
}

type ServiceExposes []ServiceExpose

type ServiceExpose struct {
	// port on the container
	Port uint32 `protobuf:"varint,1,opt,name=port,proto3" json:"port" yaml:"port"`
	// port on the service definition
	ExternalPort uint32                   `protobuf:"varint,2,opt,name=external_port,json=externalPort,proto3" json:"externalPort" yaml:"externalPort"`
	Proto        ServiceProtocol          `protobuf:"bytes,3,opt,name=proto,proto3,casttype=ServiceProtocol" json:"proto" yaml:"proto"`
	Service      string                   `protobuf:"bytes,4,opt,name=service,proto3" json:"service" yaml:"service"`
	Global       bool                     `protobuf:"varint,5,opt,name=global,proto3" json:"global" yaml:"global"`
	Hosts        []string                 `protobuf:"bytes,6,rep,name=hosts,proto3" json:"hosts" yaml:"hosts"`
	HTTPOptions  ServiceExposeHTTPOptions `protobuf:"bytes,7,opt,name=http_options,json=httpOptions,proto3" json:"httpOptions" yaml:"httpOptions"`
	// The name of the IP address associated with this, if any
	IP string `protobuf:"bytes,8,opt,name=ip,proto3" json:"ip" yaml:"ip"`
	// The sequence number of the associated endpoint in the on-chain data
	EndpointSequenceNumber uint32 `protobuf:"varint,9,opt,name=endpoint_sequence_number,json=endpointSequenceNumber,proto3" json:"endpointSequenceNumber" yaml:"endpointSequenceNumber"`
}

type ServiceProtocol string

const (
	TCP = ServiceProtocol("TCP")
	UDP = ServiceProtocol("UDP")
)

type ServiceExposeHTTPOptions struct {
	MaxBodySize uint32   `protobuf:"varint,1,opt,name=max_body_size,json=maxBodySize,proto3" json:"maxBodySize" yaml:"maxBodySize"`
	ReadTimeout uint32   `protobuf:"varint,2,opt,name=read_timeout,json=readTimeout,proto3" json:"readTimeout" yaml:"readTimeout"`
	SendTimeout uint32   `protobuf:"varint,3,opt,name=send_timeout,json=sendTimeout,proto3" json:"sendTimeout" yaml:"sendTimeout"`
	NextTries   uint32   `protobuf:"varint,4,opt,name=next_tries,json=nextTries,proto3" json:"nextTries" yaml:"nextTries"`
	NextTimeout uint32   `protobuf:"varint,5,opt,name=next_timeout,json=nextTimeout,proto3" json:"nextTimeout" yaml:"nextTimeout"`
	NextCases   []string `protobuf:"bytes,6,rep,name=next_cases,json=nextCases,proto3" json:"nextCases" yaml:"nextCases"`
}

// ServiceParams
type ServiceParams struct {
	Storage []StorageParams `protobuf:"bytes,1,rep,name=storage,proto3" json:"storage" yaml:"storage"`
}

type StorageParams struct {
	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" yaml:"name"`
	Mount    string `protobuf:"bytes,2,opt,name=mount,proto3" json:"mount" yaml:"mount"`
	ReadOnly bool   `protobuf:"varint,3,opt,name=read_only,json=readOnly,proto3" json:"readOnly" yaml:"readOnly"`
}

type ServiceImageCredentials struct {
	Host     string `protobuf:"bytes,1,opt,name=host,proto3" json:"host" yaml:"host"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email" yaml:"email"`
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username" yaml:"username"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password" yaml:"password"`
}
