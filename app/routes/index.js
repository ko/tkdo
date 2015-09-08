import Ember from 'ember';

export default Ember.Route.extend({

    model: function() {
        return {
            data: this.store.findAll('task'),
            task: {}
        };
    },

    actions: {
        createTask(info) {

            let newTask = this.store.createRecord('task', {
                title: info.title,
                description: info.description,
                dateCreated: new Date(),
                dateCompleted: null,
                author: "FIXME",
            });

            newTask.save();
        },

        completeTask(info) {
            this.store.find('task', info.id).then(function (task) {
                Ember.set(task, "dateCompleted", new Date());

                // Don't delete; just don't display..
                // the show-task-list will check for non-null 
                // value in task.dateCompleted
                //
                //   NO: task.destroyRecord();  
                //
                task.save();
            });
        },
    }
});
