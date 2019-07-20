devinp@ip-172-16-2-217:/data$ qstat -f -xml
<?xml version='1.0'?>
<job_info  xmlns:xsd="http://arc.liv.ac.uk/repos/darcs/sge/source/dist/util/resources/schemas/qstat/qstat.xsd">
  <queue_info>
    <Queue-List>
      <name>all.q@ip-172-16-2-146.us-west-2.compute.internal</name>
      <qtype>BIP</qtype>
      <slots_used>0</slots_used>
      <slots_resv>0</slots_resv>
      <slots_total>8</slots_total>
      <load_avg>0.01000</load_avg>
      <arch>lx-amd64</arch>
    </Queue-List>
    <Queue-List>
      <name>all.q@ip-172-16-2-70.us-west-2.compute.internal</name>
      <qtype>BIP</qtype>
      <slots_used>0</slots_used>
      <slots_resv>0</slots_resv>
      <slots_total>8</slots_total>
      <load_avg>0.04000</load_avg>
      <arch>lx-amd64</arch>
    </Queue-List>
  </queue_info>
  <job_info>
  </job_info>
</job_info>
devinp@ip-172-16-2-217:/data$ qsub -b y sleep 10 && date
Your job 1 ("sleep") has been submitted
Mon Jul 15 19:27:24 UTC 2019
devinp@ip-172-16-2-217:/data$ qstat -f -xml
<?xml version='1.0'?>
<job_info  xmlns:xsd="http://arc.liv.ac.uk/repos/darcs/sge/source/dist/util/resources/schemas/qstat/qstat.xsd">
  <queue_info>
    <Queue-List>
      <name>all.q@ip-172-16-2-146.us-west-2.compute.internal</name>
      <qtype>BIP</qtype>
      <slots_used>1</slots_used>
      <slots_resv>0</slots_resv>
      <slots_total>8</slots_total>
      <load_avg>0.01000</load_avg>
      <arch>lx-amd64</arch>
      <job_list state="running">
        <JB_job_number>1</JB_job_number>
        <JAT_prio>0.55500</JAT_prio>
        <JB_name>sleep</JB_name>
        <JB_owner>devinp</JB_owner>
        <state>r</state>
        <JAT_start_time>2019-07-15T19:27:25</JAT_start_time>
        <slots>1</slots>
      </job_list>
    </Queue-List>
    <Queue-List>
      <name>all.q@ip-172-16-2-70.us-west-2.compute.internal</name>
      <qtype>BIP</qtype>
      <slots_used>0</slots_used>
      <slots_resv>0</slots_resv>
      <slots_total>8</slots_total>
      <load_avg>0.04000</load_avg>
      <arch>lx-amd64</arch>
    </Queue-List>
  </queue_info>
  <job_info>
  </job_info>
</job_info>
devinp@ip-172-16-2-217:/data$ qstat -f -xml
<?xml version='1.0'?>
<job_info  xmlns:xsd="http://arc.liv.ac.uk/repos/darcs/sge/source/dist/util/resources/schemas/qstat/qstat.xsd">
  <queue_info>
    <Queue-List>
      <name>all.q@ip-172-16-2-146.us-west-2.compute.internal</name>
      <qtype>BIP</qtype>
      <slots_used>1</slots_used>
      <slots_resv>0</slots_resv>
      <slots_total>8</slots_total>
      <load_avg>0.01000</load_avg>
      <arch>lx-amd64</arch>
      <job_list state="running">
        <JB_job_number>1</JB_job_number>
        <JAT_prio>0.55500</JAT_prio>
        <JB_name>sleep</JB_name>
        <JB_owner>devinp</JB_owner>
        <state>r</state>
        <JAT_start_time>2019-07-15T19:27:25</JAT_start_time>
        <slots>1</slots>
      </job_list>
    </Queue-List>
    <Queue-List>
      <name>all.q@ip-172-16-2-70.us-west-2.compute.internal</name>
      <qtype>BIP</qtype>
      <slots_used>0</slots_used>
      <slots_resv>0</slots_resv>
      <slots_total>8</slots_total>
      <load_avg>0.04000</load_avg>
      <arch>lx-amd64</arch>
    </Queue-List>
  </queue_info>
  <job_info>
  </job_info>
</job_info>
devinp@ip-172-16-2-217:/data$ qsub -b y sleep 10 && date
Your job 2 ("sleep") has been submitted
Mon Jul 15 19:27:54 UTC 2019
devinp@ip-172-16-2-217:/data$ qsub -b y sleep 10 && date
Your job 3 ("sleep") has been submitted
Mon Jul 15 19:27:55 UTC 2019
devinp@ip-172-16-2-217:/data$ qsub -b y sleep 10 && date
Your job 4 ("sleep") has been submitted
Mon Jul 15 19:27:55 UTC 2019
devinp@ip-172-16-2-217:/data$ qsub -b y sleep 10 && date
Your job 5 ("sleep") has been submitted
Mon Jul 15 19:27:56 UTC 2019
devinp@ip-172-16-2-217:/data$ qsub -b y sleep 10 && date
Your job 6 ("sleep") has been submitted
Mon Jul 15 19:27:56 UTC 2019
devinp@ip-172-16-2-217:/data$ qstat -f -xml
<?xml version='1.0'?>
<job_info  xmlns:xsd="http://arc.liv.ac.uk/repos/darcs/sge/source/dist/util/resources/schemas/qstat/qstat.xsd">
  <queue_info>
    <Queue-List>
      <name>all.q@ip-172-16-2-146.us-west-2.compute.internal</name>
      <qtype>BIP</qtype>
      <slots_used>1</slots_used>
      <slots_resv>0</slots_resv>
      <slots_total>8</slots_total>
      <load_avg>0.01000</load_avg>
      <arch>lx-amd64</arch>
      <job_list state="running">
        <JB_job_number>2</JB_job_number>
        <JAT_prio>0.55500</JAT_prio>
        <JB_name>sleep</JB_name>
        <JB_owner>devinp</JB_owner>
        <state>r</state>
        <JAT_start_time>2019-07-15T19:27:55</JAT_start_time>
        <slots>1</slots>
      </job_list>
    </Queue-List>
    <Queue-List>
      <name>all.q@ip-172-16-2-70.us-west-2.compute.internal</name>
      <qtype>BIP</qtype>
      <slots_used>0</slots_used>
      <slots_resv>0</slots_resv>
      <slots_total>8</slots_total>
      <load_avg>0.05000</load_avg>
      <arch>lx-amd64</arch>
    </Queue-List>
  </queue_info>
  <job_info>
    <job_list state="pending">
      <JB_job_number>3</JB_job_number>
      <JAT_prio>0.00000</JAT_prio>
      <JB_name>sleep</JB_name>
      <JB_owner>devinp</JB_owner>
      <state>qw</state>
      <JB_submission_time>2019-07-15T19:27:55</JB_submission_time>
      <slots>1</slots>
    </job_list>
    <job_list state="pending">
      <JB_job_number>4</JB_job_number>
      <JAT_prio>0.00000</JAT_prio>
      <JB_name>sleep</JB_name>
      <JB_owner>devinp</JB_owner>
      <state>qw</state>
      <JB_submission_time>2019-07-15T19:27:55</JB_submission_time>
      <slots>1</slots>
    </job_list>
    <job_list state="pending">
      <JB_job_number>5</JB_job_number>
      <JAT_prio>0.00000</JAT_prio>
      <JB_name>sleep</JB_name>
      <JB_owner>devinp</JB_owner>
      <state>qw</state>
      <JB_submission_time>2019-07-15T19:27:56</JB_submission_time>
      <slots>1</slots>
    </job_list>
    <job_list state="pending">
      <JB_job_number>6</JB_job_number>
      <JAT_prio>0.00000</JAT_prio>
      <JB_name>sleep</JB_name>
      <JB_owner>devinp</JB_owner>
      <state>qw</state>
      <JB_submission_time>2019-07-15T19:27:56</JB_submission_time>
      <slots>1</slots>
    </job_list>
  </job_info>
</job_info>
devinp@ip-172-16-2-217:/data$ qstat -f -xml
<?xml version='1.0'?>
<job_info  xmlns:xsd="http://arc.liv.ac.uk/repos/darcs/sge/source/dist/util/resources/schemas/qstat/qstat.xsd">
  <queue_info>
    <Queue-List>
      <name>all.q@ip-172-16-2-146.us-west-2.compute.internal</name>
      <qtype>BIP</qtype>
      <slots_used>0</slots_used>
      <slots_resv>0</slots_resv>
      <slots_total>8</slots_total>
      <load_avg>0.01000</load_avg>
      <arch>lx-amd64</arch>
    </Queue-List>
    <Queue-List>
      <name>all.q@ip-172-16-2-70.us-west-2.compute.internal</name>
      <qtype>BIP</qtype>
      <slots_used>0</slots_used>
      <slots_resv>0</slots_resv>
      <slots_total>8</slots_total>
      <load_avg>0.05000</load_avg>
      <arch>lx-amd64</arch>
    </Queue-List>
  </queue_info>
  <job_info>
    <job_list state="pending">
      <JB_job_number>3</JB_job_number>
      <JAT_prio>0.00000</JAT_prio>
      <JB_name>sleep</JB_name>
      <JB_owner>devinp</JB_owner>
      <state>qw</state>
      <JB_submission_time>2019-07-15T19:27:55</JB_submission_time>
      <slots>1</slots>
    </job_list>
    <job_list state="pending">
      <JB_job_number>4</JB_job_number>
      <JAT_prio>0.00000</JAT_prio>
      <JB_name>sleep</JB_name>
      <JB_owner>devinp</JB_owner>
      <state>qw</state>
      <JB_submission_time>2019-07-15T19:27:55</JB_submission_time>
      <slots>1</slots>
    </job_list>
    <job_list state="pending">
      <JB_job_number>5</JB_job_number>
      <JAT_prio>0.00000</JAT_prio>
      <JB_name>sleep</JB_name>
      <JB_owner>devinp</JB_owner>
      <state>qw</state>
      <JB_submission_time>2019-07-15T19:27:56</JB_submission_time>
      <slots>1</slots>
    </job_list>
    <job_list state="pending">
      <JB_job_number>6</JB_job_number>
      <JAT_prio>0.00000</JAT_prio>
      <JB_name>sleep</JB_name>
      <JB_owner>devinp</JB_owner>
      <state>qw</state>
      <JB_submission_time>2019-07-15T19:27:56</JB_submission_time>
      <slots>1</slots>
    </job_list>
  </job_info>
</job_info>
devinp@ip-172-16-2-217:/data$ qstat -f -xml
<?xml version='1.0'?>
<job_info  xmlns:xsd="http://arc.liv.ac.uk/repos/darcs/sge/source/dist/util/resources/schemas/qstat/qstat.xsd">
  <queue_info>
    <Queue-List>
      <name>all.q@ip-172-16-2-146.us-west-2.compute.internal</name>
      <qtype>BIP</qtype>
      <slots_used>0</slots_used>
      <slots_resv>0</slots_resv>
      <slots_total>8</slots_total>
      <load_avg>0.01000</load_avg>
      <arch>lx-amd64</arch>
    </Queue-List>
    <Queue-List>
      <name>all.q@ip-172-16-2-70.us-west-2.compute.internal</name>
      <qtype>BIP</qtype>
      <slots_used>0</slots_used>
      <slots_resv>0</slots_resv>
      <slots_total>8</slots_total>
      <load_avg>0.05000</load_avg>
      <arch>lx-amd64</arch>
    </Queue-List>
  </queue_info>
  <job_info>
    <job_list state="pending">
      <JB_job_number>3</JB_job_number>
      <JAT_prio>0.00000</JAT_prio>
      <JB_name>sleep</JB_name>
      <JB_owner>devinp</JB_owner>
      <state>qw</state>
      <JB_submission_time>2019-07-15T19:27:55</JB_submission_time>
      <slots>1</slots>
    </job_list>
    <job_list state="pending">
      <JB_job_number>4</JB_job_number>
      <JAT_prio>0.00000</JAT_prio>
      <JB_name>sleep</JB_name>
      <JB_owner>devinp</JB_owner>
      <state>qw</state>
      <JB_submission_time>2019-07-15T19:27:55</JB_submission_time>
      <slots>1</slots>
    </job_list>
    <job_list state="pending">
      <JB_job_number>5</JB_job_number>
      <JAT_prio>0.00000</JAT_prio>
      <JB_name>sleep</JB_name>
      <JB_owner>devinp</JB_owner>
      <state>qw</state>
      <JB_submission_time>2019-07-15T19:27:56</JB_submission_time>
      <slots>1</slots>
    </job_list>
    <job_list state="pending">
      <JB_job_number>6</JB_job_number>
      <JAT_prio>0.00000</JAT_prio>
      <JB_name>sleep</JB_name>
      <JB_owner>devinp</JB_owner>
      <state>qw</state>
      <JB_submission_time>2019-07-15T19:27:56</JB_submission_time>
      <slots>1</slots>
    </job_list>
  </job_info>
</job_info>
devinp@ip-172-16-2-217:/data$ qstat -f -xml
<?xml version='1.0'?>
<job_info  xmlns:xsd="http://arc.liv.ac.uk/repos/darcs/sge/source/dist/util/resources/schemas/qstat/qstat.xsd">
  <queue_info>
    <Queue-List>
      <name>all.q@ip-172-16-2-146.us-west-2.compute.internal</name>
      <qtype>BIP</qtype>
      <slots_used>2</slots_used>
      <slots_resv>0</slots_resv>
      <slots_total>8</slots_total>
      <load_avg>0.01000</load_avg>
      <arch>lx-amd64</arch>
      <job_list state="running">
        <JB_job_number>3</JB_job_number>
        <JAT_prio>0.55500</JAT_prio>
        <JB_name>sleep</JB_name>
        <JB_owner>devinp</JB_owner>
        <state>r</state>
        <JAT_start_time>2019-07-15T19:28:10</JAT_start_time>
        <slots>1</slots>
      </job_list>
      <job_list state="running">
        <JB_job_number>5</JB_job_number>
        <JAT_prio>0.55500</JAT_prio>
        <JB_name>sleep</JB_name>
        <JB_owner>devinp</JB_owner>
        <state>r</state>
        <JAT_start_time>2019-07-15T19:28:10</JAT_start_time>
        <slots>1</slots>
      </job_list>
    </Queue-List>
    <Queue-List>
      <name>all.q@ip-172-16-2-70.us-west-2.compute.internal</name>
      <qtype>BIP</qtype>
      <slots_used>2</slots_used>
      <slots_resv>0</slots_resv>
      <slots_total>8</slots_total>
      <load_avg>0.04000</load_avg>
      <arch>lx-amd64</arch>
      <job_list state="running">
        <JB_job_number>4</JB_job_number>
        <JAT_prio>0.55500</JAT_prio>
        <JB_name>sleep</JB_name>
        <JB_owner>devinp</JB_owner>
        <state>r</state>
        <JAT_start_time>2019-07-15T19:28:10</JAT_start_time>
        <slots>1</slots>
      </job_list>
      <job_list state="running">
        <JB_job_number>6</JB_job_number>
        <JAT_prio>0.55500</JAT_prio>
        <JB_name>sleep</JB_name>
        <JB_owner>devinp</JB_owner>
        <state>r</state>
        <JAT_start_time>2019-07-15T19:28:10</JAT_start_time>
        <slots>1</slots>
      </job_list>
    </Queue-List>
  </queue_info>
  <job_info>
  </job_info>
</job_info>
devinp@ip-172-16-2-217:/data$ qstat -f -xml
<?xml version='1.0'?>
<job_info  xmlns:xsd="http://arc.liv.ac.uk/repos/darcs/sge/source/dist/util/resources/schemas/qstat/qstat.xsd">
  <queue_info>
    <Queue-List>
      <name>all.q@ip-172-16-2-146.us-west-2.compute.internal</name>
      <qtype>BIP</qtype>
      <slots_used>0</slots_used>
      <slots_resv>0</slots_resv>
      <slots_total>8</slots_total>
      <load_avg>0.01000</load_avg>
      <arch>lx-amd64</arch>
    </Queue-List>
    <Queue-List>
      <name>all.q@ip-172-16-2-70.us-west-2.compute.internal</name>
      <qtype>BIP</qtype>
      <slots_used>0</slots_used>
      <slots_resv>0</slots_resv>
      <slots_total>8</slots_total>
      <load_avg>0.04000</load_avg>
      <arch>lx-amd64</arch>
    </Queue-List>
  </queue_info>
  <job_info>
  </job_info>
</job_info>