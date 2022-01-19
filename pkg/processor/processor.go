package processor

import (
	"encoding/json"
	"io"
)

// StepFunc is the signature of the function used to manipulate
// JSON data read by Processor
type StepFunc func(*map[string]interface{}) (bool, error)

// Processor as the name implies processes JSON by reading from
// decoder and writing to encoder.
// All data read will be manipulated sequentially by the steps.
type Processor struct {
	decoder *json.Decoder
	encoder *json.Encoder
	steps   []StepFunc
}

// New creates a Processor that reads from r and writes to w
func New(r io.Reader, w io.Writer) *Processor {
	p := &Processor{
		decoder: json.NewDecoder(r),
		encoder: json.NewEncoder(w),
	}
	return p
}

// AddStep appends a new step function to p.steps.
func (p *Processor) AddStep(step StepFunc) {
	p.steps = append(p.steps, step)
}

// Next reads the next JSON encoded value from p.decoder
// and writes it to p.encoder after applying all the
// the p.steps functions.
// If any step function returns false the value is skipped.
// Returns error if it fails reading from decoder, writing to
// encoder or if any step function fails.
// Return io.EOF if no more data can be read from p.decoder.
func (p *Processor) Next() error {
	var value map[string]interface{}
	err := p.decoder.Decode(&value)
	if err != nil {
		return err
	}

	for _, step := range p.steps {
		if accept, err := step(&value); !accept {
			return nil
		} else if err != nil {
			return err
		}
	}

	err = p.encoder.Encode(value)
	if err != nil {
		return err
	}

	return nil
}

// ReadAll reads all data from p.decoder, it only
// stops if there is nothing more to read.
// Returns error if p.Next() fails for any reason.
func (p *Processor) ReadAll() error {
	for {
		err := p.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}
	return nil
}
