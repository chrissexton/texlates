\documentclass[letter]{article}

\usepackage[utf8]{inputenc}
\usepackage[letterpaper, margin=1cm]{geometry}
\usepackage{array}
\usepackage{colortbl}
\usepackage[x11names]{xcolor}

\pagenumbering{gobble}

%\renewcommand{\arraystretch}{1.17}
\renewcommand{\arraystretch}{1.6}

\definecolor{anti-flashwhite}{rgb}{0.95, 0.95, 0.96}

\newcommand*{\grayline}{%
	\arrayrulecolor{anti-flashwhite}
	\cline{1-3}
	\arrayrulecolor{black}
}

% TODO: Make this reactive
\newcommand\textbox[1]{%
	\parbox{6cm}{#1}%
}

\begin{document}
	\centering
	
