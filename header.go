package cat

type Header interface {
	Encodable
	GetDomain() string
	GetHostname() string
	GetIpAddress() string
}

type header struct {
	m_domain    string
	m_hostname  string
	m_ipAddress string
}

func NewHeader() Header {

	return header{
		DOMAIN,
		HOSTNAME,
		IP,
	}
}

func (h header) GetDomain() string {
	return h.m_domain
}

func (h header) GetHostname() string {
	return h.m_hostname
}

func (h header) GetIpAddress() string {
	return h.m_ipAddress
}
