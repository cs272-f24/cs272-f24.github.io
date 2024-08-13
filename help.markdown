---
layout: default
title: Getting Help
permalink: /help/
---

### CampusWire

1. All registered students have been sent an invitation to the CampusWire forum via your dons email address
1. Etiquette
    1. Please post **technical questions** in the Class Feed and answer each others' questions so CampusWire is useful for all.
    1. You may post anonymously if you must, but you're only anonymous to other students, not the instructor or TAs.
    1. Please post **grading questions** to Instructors and TAs.
    1. We will use unresolved questions as a to-do list. If our reply doesn't answer your question, please un-resolve it so we can see your followup.
    1. Please **do not DM the instructor or TAs** unless you really have to. DMs in CampusWire are not a good to-do list. Also 1:1 messages are likely to take longer to answer due to teaching/commuting time during the week.

### Online Resources

1. [Official Go documentation](https://go.dev/doc/)
1. [Go By Example](https://gobyexample.com/)
1. [GoBridge](http://golangbridge.org/)
1. [Women Who Go](https://womenwhogo.org/)

### Drop-In Office Hours

{% for person in site.people -%}
| [{{person.name}}](mailto:{{person.email}}) ({{person.role}}) | {{person.office_hours}} |
{% endfor %}

If our regular office hours don't match your schedule, please contact us on CampusWire to make an appointment at a mutually available time.

### CS Tutoring Center

The [Computer Science Tutoring Center](https://tutoringcenter.cs.usfca.edu/) is a great resource to get help from other students.
