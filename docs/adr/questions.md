1. Actual production use (daily note user) don't need to install cli so need
   a package that will handle migration lifecyle

2. Where to store migration multiple places understand when to use which ?
   i. source file will be embeded to the binary golang generates.

   Choice by process of elimination

   ii. remote repository

   Only reason would internet issue, so new release would have new golang code
   but not the database changes reflected
   [a big if user does not have internet how they got the new release but i digress]

   iii. embeded database file to be copied on user database

   Not good, if any user update tnotes the database update will not be easy to provide
   [just imagine if user use our application to upgrade for feature]

3. should there be a need of a custom that migration files are present but
   only applied to specific migration file why ?

4. There is no such thing as secret env in a build its a binary so it can be reversed engineered
