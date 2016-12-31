package converter

type DefaultConverter ConversionPattern

func (c *DefaultConverter) Convert() (string, error) {
	return c.Text, nil
}
