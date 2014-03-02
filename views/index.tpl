<h1>Blog posts</h1>

<ul>
{{range .blogs}}
   <li>
      <a href="/view/{{.Id}}">{{.Title}}</a>
      from  {{.Created}}
      <a href="/edit/{{.Id}}">Eidt</a>
      <a href="/delete/{{.Id}}">Delete</a>
   </li>
{{end}}
</ul>