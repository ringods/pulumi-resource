FROM pulumi/pulumi:v3.13.2

COPY check /opt/resource/check
COPY in /opt/resource/in
COPY out /opt/resource/out

ENTRYPOINT ["/bin/sh"]