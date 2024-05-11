package usecase

import (
	"context"
	"fmt"
	"go-boiler-plate/internal/users/model"
	"go-boiler-plate/pkg/config"
	"time"
)

type ServiceOption struct {
	Config config.AppConfig
}

type ServiceImplementation struct {
	config config.AppConfig
}

func NewServiceImplementation(opts ServiceOption) *ServiceImplementation {
	return &ServiceImplementation{
		config: opts.Config,
	}
}

func (s ServiceImplementation) Greeting(ctx context.Context, request model.GreetingRequest) (result model.GreetingResponse, err error) {
	err = request.Validate()
	if err != nil {
		return
	}
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return
	}
	birthDate := time.Date(request.Year, time.Month(request.Month), request.Date, 0, 0, 0, 0, loc)
	now := time.Now()
	if now.Before(birthDate) {
		return
	}

	result.Name = request.Name
	result.AgeYears, result.AgeMonths, result.AgeDays = s.calculateAge(birthDate)
	result.Zodiac = s.getZodiacSign(birthDate.Month(), birthDate.Day())
	result.Text = fmt.Sprintf("Hallo [%s], \nUsia anda saat ini adalah: [%d] tahun \n[%d] bulan\n[%d] Hari\n\n Bintang Anda Adalah [%s]",
		result.Name, result.AgeYears, result.AgeMonths, result.AgeDays, result.Zodiac)
	fmt.Println(result.Text)
	result.Env = s.config.App.Env
	return
}

func (s ServiceImplementation) calculateAge(birthDate time.Time) (years, months, days int) {
	now := time.Now()
	years = now.Year() - birthDate.Year()
	if now.Month() < birthDate.Month() || (now.Month() == birthDate.Month() && now.Day() < birthDate.Day()) {
		years--
	}
	months = int(now.Month()) - int(birthDate.Month())
	if months < 0 {
		months += 12
	}

	if now.Day() < birthDate.Day() {
		months--
		daysInLastMonth := time.Date(now.Year(), now.Month()-1, 0, 0, 0, 0, 0, time.UTC).Day()
		days = daysInLastMonth + now.Day() - birthDate.Day()
	} else {
		days = now.Day() - birthDate.Day()
	}

	return years, months, days
}

func (s ServiceImplementation) getZodiacSign(month time.Month, day int) string {
	switch month {
	case time.January:
		if day < 20 {
			return "Capricorn"
		}
		return "Aquarius"
	case time.February:
		if day < 19 {
			return "Aquarius"
		}
		return "Pisces"
	case time.March:
		if day < 21 {
			return "Pisces"
		}
		return "Aries"
	case time.April:
		if day < 20 {
			return "Aries"
		}
		return "Taurus"
	case time.May:
		if day < 21 {
			return "Taurus"
		}
		return "Gemini"
	case time.June:
		if day < 21 {
			return "Gemini"
		}
		return "Cancer"
	case time.July:
		if day < 23 {
			return "Cancer"
		}
		return "Leo"
	case time.August:
		if day < 23 {
			return "Leo"
		}
		return "Virgo"
	case time.September:
		if day < 23 {
			return "Virgo"
		}
		return "Libra"
	case time.October:
		if day < 23 {
			return "Libra"
		}
		return "Scorpio"
	case time.November:
		if day < 22 {
			return "Scorpio"
		}
		return "Sagittarius"
	default:
		return "Capricorn"
	}
}
