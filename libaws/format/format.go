package format

type (
	// Root ...
	Root struct {
		OutDir            string     `yaml:"OutDir"`
		Templates         []Template `yaml:"Templates"`
		TemplateToExecute string     `yaml:"TemplateToExecute"`
		Datas             []string   `yaml:"Datas"`
		Lambdas           []Lambda   `yaml:"Lambdas"`
	}
	// Template ...
	Template struct {
		Path    string `yaml:"Path"`
		Name    string `yaml:"Name"`
		OutDir  string `yaml:"OutDir"`
		OutFile string `yaml:"OutFile"`
	}
	// Lambda ...
	Lambda struct {
		Path    string   `yaml:"Path"`
		Ignore  []string `yaml:"Ignore"`
		OutDir  string   `yaml:"OutDir"`
		OutFile string   `yaml:"OutFile"`
	}
)
