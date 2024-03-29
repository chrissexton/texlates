\documentclass[a5paper]{article}
%\input \jobname

\def \dayOfWeek{Wednesday}
\def \date{2019/08/07}

\usepackage[utf8]{inputenc}
\usepackage[a5paper, margin=1cm]{geometry}
\usepackage{array}
\usepackage{colortbl}
\usepackage[x11names]{xcolor}

\pagenumbering{gobble}

\renewcommand{\arraystretch}{1.17}

\definecolor{anti-flashwhite}{rgb}{0.95, 0.95, 0.96}

\newcommand*{\grayline}{%
	\arrayrulecolor{anti-flashwhite}
	\cline{1-3}
	\arrayrulecolor{black}
}

% TODO: Make this reactive
\newcommand\textbox[1]{%
	\parbox{3.5cm}{#1}%
}

\begin{document}
	\centering
	
\begin{tabular}{ |r|r|p{35mm}|p{48mm}| }
		\hline
		\multicolumn{4}{|c|}{ 
			\cellcolor{anti-flashwhite}
			\noindent
				\textbox{\date \hfill}
				\textbox{\hfil Daily Planner \hfil}
				\textbox{\hfill \dayOfWeek}
		}\\
		\hline
		\hline
		\multicolumn{1}{|c|}{Time} & 
		\multicolumn{1}{ c|}{i.} &
		\multicolumn{1}{ c|}{Schedule} & 
		\multicolumn{1}{ c|}{Bullets}
		\\\hline
		
		6:00 &&&\\\grayline
		6:30 &&&\\\grayline
		7:00 &&&\\\grayline
		7:30 &&&\\\grayline
		8:00 &&&\\\grayline
		8:30 &&&\\\grayline
		9:00 &&&\\\grayline
		9:30 &&&\\\grayline
		10:00 &&&\\\grayline
		10:30 &&&\\\grayline
		11:00 &&&\\\grayline
		11:30 &&&\\\grayline
		12:00 &&&\\\grayline
		12:30 &&&\\\grayline
		13:00 &&&\\\grayline
		13:30 &&&\\\grayline
		14:00 &&&\\\grayline
		14:30 &&&\\\grayline
		15:00 &&&\\\grayline
		15:30 &&&\\\grayline
		16:00 &&&\\\grayline
		16:30 &&&\\\grayline
		17:00 &&&\\\grayline
		17:30 &&&\\\grayline
		18:00 &&&\\\grayline
		18:30 &&&\\\grayline
		19:00 &&&\\\grayline
		19:30 &&&\\\grayline
		20:00 &&&\\\grayline
		20:30 &&&\\\grayline
		21:00 &&&\\\grayline
		21:30 &&&\\\grayline
		22:00 &&&\\\grayline
		22:30 &&&\\\grayline
		23:00 &&&\\\grayline
		23:30 &&&\\
		
		\hline
		
	\end{tabular}
	
\end{document}
