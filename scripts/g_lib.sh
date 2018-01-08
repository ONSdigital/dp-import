#!/bin/false dotme

g_colr() { local col=$1 recurse= rst=$'\e[0m'; shift  # '[-r] bright_white_on_red' WHITE_on_red bold -- all valid
    if [[ $col == -r ]]; then recurse=1; col=$1; shift; fi
    if [[ -t 0 ]]; then
        col=$'\e['$(perl -E '$_="0;";$c=shift @ARGV;%c=(bold=>1,black=>30,red=>31,green=>32,yellow=>33,blue=>34,magenta=>35,cyan=>36,white=>37,reset=>0);$_="1;" if $c =~ /^[A-Z]/ and $c=lc $c or $c =~ s/^bright_//;$_.=10+$c{$1}.";" if $c =~ s/_on_(\w+)$// and defined $c{$1};if(defined $c{$c}){$_.=$c{$c}}else{$_=$c}say' $col)m
        if [[ -n $recurse ]]; then set -- "${@//$rst/$col}"; fi
        echo -e "$col""$@""$rst"
    else
        echo "$@"
    fi
}
g_ts()      { echo  $(g_colr    BLUE   $(date '+%F %T')) ${g_ts_host:+$(g_colr green $g_ts_host)} "$@"; }
g_info()    { g_ts "$(g_colr -r cyan   "$@")"; }
g_trace()   { g_info "$@" >&2; }
g_info2()   { g_info "$@" >&2; } # DEPRECATED
g_err()     { g_ts "$(g_colr -r RED    ERROR "$@")" >&2; }
g_warn()    { g_ts "$(g_colr -r YELLOW WARN  "$@")" >&2; }
g_log()     { g_ts "$@"; }
g_opts()    { while [[ -n $1 ]]; do if [[ $1 == host ]]; then g_ts_host=$myHOST; fi; shift; done; }
g_zsh()     { [[ -n $ZSH_NAME ]]; }
g_row_col() { local X= R= C=; if g_zsh; then IFS=';[' read -sdR X\?$'\E[6n' R C; else IFS=';[' read -sdR -p $'\E[6n' X R C; fi; echo $R $C; }
g_col()     { local rc="$(g_row_col)"; echo ${rc#* }; }
g_cont()    { local res=$1 arg=--no; shift; if [[ $res == -y ]]; then arg=; res=$1; shift; fi; yorn $arg "$@" Continue after $(g_colr bright_white_on_red error code $res) || exit $res; }

g_die() {
    local res=$1; shift
    g_err "$@"
    exit $res
}

g_ensure_env() { local e=;for e; do [[ -z ${!e} ]] && g_die 2 $e unset; done; }
g_ensure_dir() { local d=;for d; do [[ -d $d ]] || mkdir -p $d || g_die 2 $d failed; done; }
g_ensure_in_path() { # [ --end ] path...
	local end_ok= p=
	while true; do
            case $1 in
                --end) end_ok=1; ;;
                *) break; ;;
            esac
            shift
        done
	for p; do
		[[ :$PATH: == *:$p:* ]] && continue
		if [[ -n $end_ok ]]; then PATH=$PATH:$p; continue; fi
		PATH=$p:$PATH
	done
}

yorn() {
	local quit_to= def=yn g_all=a do_it= res= cont= comment= cont_ok= no_ok= any_key= hit=
	while [[ $1 == --* ]]; do
		if   [[ $1 == --           ]]; then shift; break
		elif [[ $1 == --any        ]]; then any_key=y
		elif [[ $1 == --comment    ]]; then comment=$(g_colr -r yellow "$2 "); shift
		elif [[ $1 == --cont       ]]; then cont=y
		elif [[ $1 == --no         ]]; then def=ny
		elif [[ $1 == --no-ok      ]]; then no_ok=1
		elif [[ $1 == --x          ]]; then do_it=1
		elif [[ $1 == --xc         ]]; then do_it=1; cont=y
		elif [[ $1 == --xc-ok      ]]; then do_it=1; cont=y; cont_ok=$2; shift
		elif [[ $1 == --no-all     ]]; then g_all=
		elif [[ $1 == --reset-all  ]]; then g_do_all=; [[ -z $2 ]] && return
		elif [[ $1 == --quit-to    ]]; then quit_to=$2; shift
		else break
		fi
		shift
	done
	if [[ -n $g_do_all ]]; then
		if [[ $def == ny ]]; then
			res=1	# skip --no when 'all'
		else
			res=0
			[[ -n $do_it ]] && g_info EXEC: "$comment"$(g_colr bright_white "$@")
		fi
	fi
	while [[ -z $res ]]; do
		hit="$(g_colr -r bright_blue "$(g_colr cyan ${def:0:1})${def:1}${g_all}hq")"
		[[ -n $any_key ]] && hit="$(g_colr -r bright_blue hit any key or: $(g_colr cyan hq))"
		if g_zsh; then read -k 1 yorn\?"$(g_info "$comment""$@") [$hit] "
		else           read -n 1 -p    "$(g_info "$comment""$@") [$hit] " yorn
		fi
		(( $(g_col) > 1 )) && echo
		[[ -n $any_key && :h:q: != *:$yorn:* ]] && return 0
		if [[ -z $yorn || $yorn == " " || $yorn == $'\n' ]]; then yorn=${def:0:1}; fi
		if   [[ $yorn == q                   ]]; then [[ -z $quit_to ]] && exit 0; $quit_to; exit 0
		elif [[ -n $g_all && $yorn == $g_all ]]; then g_do_all=yes; res=0
		elif [[ $yorn == h                   ]]; then g_info "Hit: $(g_colr bright_magenta y)=yes, $(g_colr bright_magenta n)=no, $(g_colr bright_magenta a)=all (accept default for all subsequent), $(g_colr bright_magenta q)=quit"
		elif [[ $yorn == n                   ]]; then res=1; if [[ -n $no_ok ]]; then do_it=; res=0; fi
		elif [[ $yorn == y                   ]]; then res=0
		fi
	done
	if [[ -n $do_it && $res == 0 ]]; then
		"$@" || res=$?
		if (( res > 0 )); then
			if [[ -n $cont ]]; then
				if [[ $cont_ok == $res ]]; then g_cont -y $res; else g_cont $res; fi
			fi
		fi
	fi
	return $res
}
