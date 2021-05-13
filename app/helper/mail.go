package helper

import (
	"fmt"

	"github.com/takemo101/go-fiber/pkg"
)

type MailTemplateInformation struct {
	Subject string
	Path    string
}

// MailTemplate mail template helper
type MailTemplate struct {
	factory   pkg.MailFactory
	Templates map[string]MailTemplateInformation
}

// NewMailTemplate constructor
func NewMailTemplate(
	factory pkg.MailFactory,
	config pkg.Config,
) MailTemplate {

	mailTemplate := MailTemplate{
		factory:   factory,
		Templates: map[string]MailTemplateInformation{},
	}

	// build mail template informations
	maps, err := config.Load("mail-template")
	if err == nil {
		for key, data := range maps {
			info := data.(map[string]interface{})
			var information MailTemplateInformation
			subject, subjectOK := info["subject"]
			if subjectOK {
				information.Subject = subject.(string)
			}

			path, pathOK := info["path"]
			if pathOK {
				information.Path = path.(string)
			}

			mailTemplate.Templates[key] = information
		}
	}

	return mailTemplate
}

// GetMailTemplateInformationByKey get mail template object by key
func (template *MailTemplate) GetMailTemplateInformationByKey(key string) (information MailTemplateInformation, err error) {
	info, ok := template.Templates[key]
	if ok {
		return info, err
	}
	return information, fmt.Errorf("mail-template information %s does not exist", key)
}

// GetMailTemplateInformationByKey get mail creator by key and data
func (template *MailTemplate) GetTextMailCreatorByKey(key string, data pkg.BindData) (creator pkg.MailCreator, err error) {
	var information MailTemplateInformation
	creator, information, err = template.createMailCreatorAndInformationByKey(key, data)
	creator.TemplateText(information.Path, data)

	return creator, err
}

// GetHTMLMailCreatorByKey get mail creator by key and data
func (template *MailTemplate) GetHTMLMailCreatorByKey(key string, data pkg.BindData) (creator pkg.MailCreator, err error) {
	var information MailTemplateInformation
	creator, information, err = template.createMailCreatorAndInformationByKey(key, data)
	creator.TemplateHTML(information.Path, data)

	return creator, err
}

// createMailCreatorAndInformationByKey create mail creator by key and data
func (template *MailTemplate) createMailCreatorAndInformationByKey(key string, data pkg.BindData) (creator pkg.MailCreator, information MailTemplateInformation, err error) {
	info, infoErr := template.GetMailTemplateInformationByKey(key)
	if infoErr != nil {
		return creator, information, infoErr
	}

	creator = template.factory.Create()
	subject, parseErr := creator.Parse(info.Subject, data)
	if parseErr != nil {
		return creator, info, parseErr
	}
	creator.Subject(subject)

	return creator, info, err
}
