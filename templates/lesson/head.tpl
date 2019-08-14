\documentclass[letter]{article}
%\input \jobname

\def \colwidth{56mm}

\usepackage[utf8]{inputenc}
\usepackage[letterpaper, margin=1cm]{geometry}
\usepackage{array}
\usepackage{graphicx}
\usepackage{colortbl}
\usepackage[x11names]{xcolor}
\usepackage{tabularx,booktabs}

\pagenumbering{gobble}

\renewcommand{\arraystretch}{1.6}

\newcolumntype{C}[1]{>{\centering\arraybackslash}p{#1}}
\newcolumntype{M}[1]{>{\centering\arraybackslash}m{#1}}
\newcolumntype{S}{>{\centering\arraybackslash} m{.4\linewidth} }

\definecolor{anti-flashwhite}{rgb}{0.95, 0.95, 0.96}

\newcommand*{\grayline}{%
	\arrayrulecolor{anti-flashwhite}
	\cline{1-3}
	\arrayrulecolor{black}
}

% TODO: Make this reactive
\newcommand\textbox[1]{%
	\parbox{60.5mm}{#1}%
}

\begin{document}
\centering


