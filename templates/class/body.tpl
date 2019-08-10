\begin{tabular}{ |r|r|p{35mm}|p{48mm}| }
		\hline
		\multicolumn{4}{|c|}{ 
			\cellcolor{anti-flashwhite}
			\noindent
				\textbox{ {{.DateRange}} \hfill}
				\textbox{\hfil Semester Planner \hfil}
				\textbox{\hfill {{.Week}} }
		}\\
		\hline
\end{tabular}

\begin{tabular}{|m{1em}|m{\colwidth}|m{\colwidth}|m{\colwidth}|@{}m{0cm}@{}}\hline
   {{range .Days}}& \centering\arraybackslash{ {{.}} } {{end}} \\\hline
   {{range .Courses}}
   \centering\rotatebox{90}{ {{.}} } &&& \\[75mm] \hline
   {{end}}
\end{tabular}
