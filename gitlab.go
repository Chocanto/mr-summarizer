package main

import (
	"bytes"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/xanzy/go-gitlab"
	"io"
	"log"
	"strconv"
	"time"
)

func gatherMRs() string {
	git := gitlab.NewClient(nil, Config.Gitlab.Token)
	git.SetBaseURL(Config.Gitlab.Url)

	projects, _, err := git.Groups.ListGroupProjects(Config.Gitlab.Group, &gitlab.ListGroupProjectsOptions{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Printf("Iterating over %d projects", len(projects))

	buf := new(bytes.Buffer)
	table := getTable(buf)

	for i := range projects {
		opts := &gitlab.ListProjectMergeRequestsOptions{
			State: gitlab.String("opened"),
		}
		MRs, _, err := git.MergeRequests.ListProjectMergeRequests(projects[i].ID, opts)
		if err != nil {
			log.Fatalln(err.Error())
		}

		if len(MRs) > 0 {
			log.Printf("Found {%d} opened MR for project {%s}\n", len(MRs), projects[i].Name)

			for _, MR := range MRs {
				if !MR.WorkInProgress {
					table.Append(formatLine(MR, projects[i]))
				}
			}
		} else {
			log.Printf("No MR found for {%s}", projects[i].Name)
		}
	}

	if table.NumLines() > 0 {
		table.Render()
		return buf.String()
	} else {
		return Printer.Sprintf(":tada: Congratulation :tada:   No merge request to validate today :parrot:")
	}
}

func getTable(writer io.Writer) *tablewriter.Table {
	table := tablewriter.NewWriter(writer)
	table.SetHeader([]string{
		Printer.Sprintf("Project"),
		Printer.Sprintf("Title"),
		Printer.Sprintf("Since"),
		":+1:",
		Printer.Sprintf("Assigned"),
	})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetAutoWrapText(false)

	return table
}

func formatLine(MR *gitlab.MergeRequest, project *gitlab.Project) []string {
	var daysS string
	diffHours := int(time.Now().Sub(*MR.CreatedAt).Hours())

	if diffHours < 24 {
		daysS = Printer.Sprintf("%dH", diffHours)
	} else {
		diffDays := int(diffHours / 24)
		daysS = Printer.Sprintf("%dD", diffDays)
		if diffDays >= Config.Threshold {
			daysS = daysS + " :warning:"
		}
	}

	projectName := fmt.Sprintf("[%.30s](%s)", project.Name, project.WebURL)
	title := fmt.Sprintf("[%.50s](%s)", MR.Title, MR.WebURL)
	var assignee string

	if MR.Assignee.ID != 0 {
		assignee = fmt.Sprintf("%s (@%s)", MR.Assignee.Name, MR.Assignee.Username)
	} else {
		assignee = Printer.Sprintf("Nobody... :sad:")
	}

	return []string{
		projectName, title, daysS, strconv.Itoa(MR.Upvotes), assignee,
	}
}
