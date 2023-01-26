package response

type Users struct {
	User_name string `json:"user_name" validate:"required,min=3,max=15"`
	Password  string `json:"password,omitempty" validate:"digit,uppercase,lowercase,special,required,min=3"`
	Email     string `json:"email"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshToken struct {
	Refresh_token string `json:"refresh_token"`
}

type Error struct {
	Message string `json:"message"`
}

type Err struct {
	Message string `json:"message"`
}

// Postgres holds postgres config
type Postgres struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	DbPort   string `yaml:"dbPort"`
	AppPort  int    `yaml:"appPort"`
	Host     string `yaml:"host"`
}

type ServerInfo struct {
	ReadTimeout       int `yaml:"read_timeout"`
	WriteTimeout      int `yaml:"write_timeout"`
	IdleTimeout       int `yaml:"idle_timeout"`
	ReadHeaderTimeout int `yaml:"read_header_timeout"`
	GracePeriod       int `yaml:"grace_period"`
}
